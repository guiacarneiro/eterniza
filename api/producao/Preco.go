package producao

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
	err := database.DB.AutoMigrate(&Preco{})
	if err != nil {
		panic("Erro criando tabela")
	}
}
