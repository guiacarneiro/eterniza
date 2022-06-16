package model

import (
	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
)

// Componente ...
type Componente struct {
	gorm.Model
	Quantidade     float64 `json:"quantidade"`
	MateriaPrimaID uint
	MateriaPrima   *MateriaPrima
	ProdutoID      uint
}

func init() {
	database.Migrate(&Componente{})
}
