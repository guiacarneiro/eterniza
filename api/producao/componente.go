package producao

import "gorm.io/gorm"

// Componente ...
type Componente struct {
	gorm.Model
	Quantidade     float64 `json:"quantidade"`
	MateriaPrimaID int
	ProdutoID      int
	MateriaPrima   MateriaPrima
}

func init() {
	db.AutoMigrate(&Componente{})
}
