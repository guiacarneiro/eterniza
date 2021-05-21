package main

import (
	"eterniza/api"
	"eterniza/api/controller"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var Router *mux.Router

	Router = mux.NewRouter()
	// Home Route
	Router.HandleFunc("/", api.SetMiddlewareJSON(controller.Home)).Methods("GET")

	// Login Route
	Router.HandleFunc("/login", api.SetMiddlewareJSON(controller.Login)).Methods("POST")

	//Users routes
	Router.HandleFunc("/users", api.SetMiddlewareJSON(controller.CreateUser)).Methods("POST")
	Router.HandleFunc("/users", api.SetMiddlewareJSON(controller.GetUsers)).Methods("GET")
	Router.HandleFunc("/users/{id}", api.SetMiddlewareJSON(controller.GetUser)).Methods("GET")
	Router.HandleFunc("/users/{id}", api.SetMiddlewareJSON(api.SetMiddlewareAuthentication(controller.UpdateUser))).Methods("PUT")
	Router.HandleFunc("/users/{id}", api.SetMiddlewareAuthentication(controller.DeleteUser)).Methods("DELETE")
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", Router))
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
	//var mat []producao.MateriaPrima
	//err := database.DB.Find(&mat).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var produto producao.Produto
	//err = database.DB.Preload("Componentes").First(&produto, 5).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//produto.Componentes[0].Quantidade = 201
	//err = database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&produto).Error
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
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var prod []producao.Produto
	//err = database.DB.Preload("Componentes").Find(&prod).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(prod)
}
