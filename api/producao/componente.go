package producao

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
	err := database.DB.AutoMigrate(&Componente{})
	if err != nil {
		panic("Erro criando tabela")
	}
}
