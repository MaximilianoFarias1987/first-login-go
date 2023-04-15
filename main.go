package main

import (
	"log"
	"login/commons"
	"login/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	//Defino variables de entorno
	os.Setenv("CONNECTION_BD", "sqlserver://sa:1234@LAPTOP-9ALNHCMO:49213?database=dbUsuariosGo")

	commons.Migrate()

	router := mux.NewRouter()

	routes.SetPersonasRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")

	log.Println(server.ListenAndServe())
}
