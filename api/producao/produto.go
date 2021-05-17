package producao

import (
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Referencia  string `gorm:"size:255;not null;unique" json:"referencia"`
	Componentes []Componente
}

func init() {
	db.AutoMigrate(&Produto{})
}
