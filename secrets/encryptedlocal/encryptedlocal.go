package encryptedlocal

import (
	"errors"

	"github.com/0xPolygon/polygon-edge/secrets"
	"github.com/0xPolygon/polygon-edge/secrets/local"
	"github.com/hashicorp/go-hclog"
)

// LocalSecretsManager is a SecretsManager that
// stores secrets encrypted locally on disk
type EncryptedLocalSecretsManager struct {
	prompt *Prompt
	logger hclog.Logger
	*local.LocalSecretsManager
	cryptHandler CryptHandler
	pwd          []byte
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

	prompt := NewPrompt("")
	cryptHandler := NewCryptHandler(
		prompt,
	)

	logger := params.Logger.Named(string(secrets.EncryptedLocal))
	// Set up the base object
	esm := &EncryptedLocalSecretsManager{
		prompt,
		logger,
		localSM,
		cryptHandler,
		nil,
	}

	return esm, esm.Setup()
}

func (esm *EncryptedLocalSecretsManager) SetSecret(name string, value []byte) error {
	esm.logger.Info("Configuring secret", "name", name)
	// hexValue := hex.EncodeToString(value)
	esm.logger.Info("Here is the raw hex value of your secret. \nPlease copy it and store it in a safe place.", name, string(value))

	confirmValue, err := esm.prompt.DefaultPrompt("Please rewrite the secret value to confirm that you have copied it down correctly.", "")
	if err != nil {
		return err
	}

	if confirmValue != string(value) {
		esm.logger.Error("The secret value you entered does not match the original value. Please try again.")
		return errors.New("secret value mismatch")
	} else {
		esm.logger.Info("The secret value you entered matches the original value. Continuing.")
	}

	if esm.pwd == nil || len(esm.pwd) == 0 {
		esm.pwd, err = esm.prompt.GeneratePassword()
		if err != nil {
			return err
		}
	}

	encryptedValue, err := esm.cryptHandler.Encrypt(value, esm.pwd)
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

	if esm.pwd == nil || len(esm.pwd) == 0 {
		esm.pwd, err = esm.prompt.InputPassword(false)
		if err != nil {
			return nil, err
		}
	}

	return esm.cryptHandler.Decrypt(encryptedValue, esm.pwd)
}

type SecretHelper interface {
	beforeSet(name string, value []byte) error
	afterSet(name string) ([]byte, error)
}

type SecretHelperFactory func() (SecretHelper, error)

var secretHandlers = map[secrets.SecretType]SecretHelperFactory{}
