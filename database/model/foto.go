package model

import (
	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
)

type Foto struct {
	gorm.Model
	url       string
	ProdutoID uint
}

func init() {
	database.Migrate(&Foto{})
}
