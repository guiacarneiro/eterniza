package producao

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
	err := database.DB.AutoMigrate(&Foto{})
	if err != nil {
		panic("Erro criando tabela")
	}
}
