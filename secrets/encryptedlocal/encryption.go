package encryptedlocal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type PasswordHandler interface {
	InputPassword() ([]byte, error)
}

type CryptHandler interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

type cryptHandler struct {
	passHandler PasswordHandler
}

func NewCryptHandler(passHandler PasswordHandler) CryptHandler {
	return &cryptHandler{
		passHandler: passHandler,
	}
}

func (ch *cryptHandler) Encrypt(data []byte) ([]byte, error) {
	key, err := ch.passHandler.InputPassword()
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func (ch *cryptHandler) Decrypt(data []byte) ([]byte, error) {
	key, err := ch.passHandler.InputPassword()
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}
