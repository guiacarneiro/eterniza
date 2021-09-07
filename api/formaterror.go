package api

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nome não disponível")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email não disponível")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Senha incorreta")
	}
	return errors.New("Dados incorretos")
}
