package routes

import (
	"login/controllers"

	"github.com/gorilla/mux"
)

func SetPersonasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/user/api").Subrouter()

	subRoute.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveUser).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetUserbyId).Methods("GET")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteUser).Methods("POST")
}
