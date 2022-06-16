package model

import (
	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
)

type TabelaPreco struct {
	gorm.Model
	Nome     string
	Aliquota float32
}

func init() {
	database.Migrate(&TabelaPreco{})
}
