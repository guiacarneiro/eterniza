package model

import (
	"github.com/guiacarneiro/eterniza/database"
)

type Foto struct {
	Base64         string
	FichaTecnicaID uint
}

func init() {
	database.Migrate(&Foto{})
}
