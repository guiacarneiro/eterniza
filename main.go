package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guiacarneiro/eterniza/api"
	"github.com/guiacarneiro/eterniza/api/controller"
	"github.com/guiacarneiro/eterniza/config"
	"log"
	"net/http"
)

func main() {
	var Router *mux.Router

	Router = mux.NewRouter()
	Router.Use(api.MiddlewareOpen)

	// Handle all preflight request
	Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	// Home Route
	Router.HandleFunc("/", controller.Home).Methods("GET")

	// Login Route
	Router.HandleFunc("/login", controller.Login).Methods("POST")

	//Users routes
	Router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	Router.HandleFunc("/users", api.SetMiddlewareAuthentication(controller.GetUsers)).Methods("GET")
	Router.HandleFunc("/users/{id}", api.SetMiddlewareAuthentication(controller.GetUser)).Methods("GET")
	Router.HandleFunc("/users/{id}", api.SetMiddlewareAuthentication(controller.UpdateUser)).Methods("PUT")
	Router.HandleFunc("/users/{id}", api.SetMiddlewareAuthentication(controller.DeleteUser)).Methods("DELETE")

	//Materia prima routes
	Router.HandleFunc("/materiaprima", api.SetMiddlewareAuthentication(controller.CreateMateriaPrima)).Methods("POST")
	Router.HandleFunc("/materiaprima", api.SetMiddlewareAuthentication(controller.GetMateriaPrimas)).Methods("GET")
	Router.HandleFunc("/materiaprima/{id}", api.SetMiddlewareAuthentication(controller.GetMateriaPrima)).Methods("GET")
	Router.HandleFunc("/materiaprima/{id}", api.SetMiddlewareAuthentication(controller.UpdateMateriaPrima)).Methods("PUT")
	Router.HandleFunc("/materiaprima/{id}", api.SetMiddlewareAuthentication(controller.DeleteMateriaPrima)).Methods("DELETE")

	//Not found router
	Router.NotFoundHandler = http.HandlerFunc(api.NotFound)

	// Start server
	server := config.GetPropriedadeDefault("server", ":8080")
	fmt.Println("Listening to " + server)
	log.Fatal(http.ListenAndServe(server, Router))
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
