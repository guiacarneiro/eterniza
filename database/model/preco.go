package model

import (
	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
)

type Preco struct {
	gorm.Model
	TabelaPrecoID uint
	TabelaPreco   *TabelaPreco
	Valor         float32
	ProdutoID     uint
}

func init() {
	database.Migrate(&Preco{})
}
