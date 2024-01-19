package encryptedlocal

import (
	"errors"

	"github.com/0xPolygon/polygon-edge/secrets"
	"github.com/0xPolygon/polygon-edge/secrets/local"
)

// LocalSecretsManager is a SecretsManager that
// stores secrets encrypted locally on disk
type EncryptedLocalSecretsManager struct {
	*local.LocalSecretsManager
	cryptHandler CryptHandler
}

// SecretsManagerFactory implements the factory method
func SecretsManagerFactory(
	_ *secrets.SecretsManagerConfig,
	params *secrets.SecretsManagerParams,
) (secrets.SecretsManager, error) {
	baseSM, err := local.SecretsManagerFactory(
		nil, // Local secrets manager doesn't require a config
		params)
	if err != nil {
		return nil, err
	}

	localSM, ok := baseSM.(*local.LocalSecretsManager)
	if !ok {
		return nil, errors.New("invalid type assertion")
	}

	prompt := NewPrompt()
	cryptHandler := NewCryptHandler(
		prompt,
	)

	// Set up the base object
	esm := &EncryptedLocalSecretsManager{
		localSM,
		cryptHandler,
	}

	return esm, esm.Setup()
}

func (esm *EncryptedLocalSecretsManager) SetSecret(name string, value []byte) error {
	encryptedValue, err := esm.cryptHandler.Encrypt(value)
	if err != nil {
		return err
	}

	return esm.LocalSecretsManager.SetSecret(name, encryptedValue)
}

func (esm *EncryptedLocalSecretsManager) GetSecret(name string) ([]byte, error) {
	encryptedValue, err := esm.LocalSecretsManager.GetSecret(name)
	if err != nil {
		return nil, err
	}

	return esm.cryptHandler.Decrypt(encryptedValue)
}
