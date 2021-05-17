package main

import (
	"eterniza/api/producao"
	"eterniza/database"
	"fmt"
	"gorm.io/gorm"
)

func main() {
	//database.DB.Save(&producao.MateriaPrima{
	//	Label: "Banho",
	//	Unity: producao.Weight,
	//	Value: 1800,
	//})
	//database.DB.Save(&producao.MateriaPrima{
	//	Label: "Strass",
	//	Unity: producao.Quantity,
	//	Value: 0.05,
	//})
	var mat []producao.MateriaPrima
	err := database.DB.Find(&mat).Error
	if err != nil {
		fmt.Println(err)
	}
	var produto producao.Produto
	err = database.DB.Preload("Componentes").First(&produto, 5).Error
	if err != nil {
		fmt.Println(err)
	}
	produto.Componentes[0].Quantidade = 201
	err = database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&produto).Error
	//err = database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&producao.Produto{
	//	Referencia: "B003",
	//	Componentes: []producao.Componente{{
	//		Quantidade:   200,
	//		MateriaPrima: mat[1],
	//	}, {
	//		Quantidade:   0.005,
	//		MateriaPrima: mat[0],
	//	}},
	//}).Error
	if err != nil {
		fmt.Println(err)
	}
	var prod []producao.Produto
	err = database.DB.Preload("Componentes").Find(&prod).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(prod)
}
