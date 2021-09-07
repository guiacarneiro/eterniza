package producao

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
	err := database.DB.AutoMigrate(&TabelaPreco{})
	if err != nil {
		panic("Erro criando tabela")
	}
}
