package routes

import (
	"login/controllers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Aquí necesitas proporcionar la clave secreta utilizada para firmar el token JWT.
			return []byte("clave-secreta"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

func SetPersonasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/user/api").Subrouter()

	subRoute.HandleFunc("/", authMiddleware(controllers.GetAllUsers)).Methods("GET")
	subRoute.HandleFunc("/save", controllers.SaveUser).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.GetUserbyId).Methods("GET")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteUser).Methods("POST")
	subRoute.HandleFunc("/login", controllers.Login).Methods("POST")
}
