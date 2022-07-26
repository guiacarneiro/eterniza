package model

import (
	"github.com/guiacarneiro/eterniza/database"
	"gorm.io/gorm"
)

// Componente ...
type Componente struct {
	gorm.Model
	Quantidade     float64       `json:"quantidade,omitempty"`
	MateriaPrimaID uint          `json:"materiaPrimaID,omitempty"`
	MateriaPrima   *MateriaPrima `json:"materiaPrima,omitempty"`
	ProdutoID      uint          `json:"produtoID,omitempty"`
}

func init() {
	database.Migrate(&Componente{})
}
