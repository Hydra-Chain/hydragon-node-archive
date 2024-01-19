package encryptedlocal

import (
	"syscall"

	"golang.org/x/term"
)

type Prompt struct {
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) InputPassword() ([]byte, error) {
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}

	return bytePassword, nil
}
