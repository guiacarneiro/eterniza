package model

import (
	"github.com/guiacarneiro/eterniza/database"
)

type Foto struct {
	url            string
	FichaTecnicaID uint
}

func init() {
	database.Migrate(&Foto{})
}
