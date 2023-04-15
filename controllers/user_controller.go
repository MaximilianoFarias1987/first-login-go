package controllers

import (
	"encoding/json"
	"log"
	"login/commons"
	"login/models"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

func GetAllUsers(writer http.ResponseWriter, reques *http.Request) {
	users := []models.User{}

	db := commons.GetConnection()

	db.Find(&users)

	json, _ := json.Marshal(users)

	commons.SendResponse(writer, http.StatusOK, json)
}

func GetUserbyId(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()

	db.Find(&user, id)

	if user.ID != uuid.Nil {
		json, _ := json.Marshal(user)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNoContent)
	}
}

func SaveUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	user.ID = uuid.New()

	log.Println(user)

	db := commons.GetConnection()

	error := json.NewDecoder(request.Body).Decode(&user)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&user).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(user)

	commons.SendResponse(writer, http.StatusCreated, json)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	// defer db.Close()

	db.Find(&user, id)

	if user.ID != uuid.Nil {
		db.Delete(user)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNoContent)
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	type Credentials struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}

	type Token struct {
		Token string `json:"token"`
	}

	user := models.User{}

	if request.Method != http.MethodPost {
		http.Error(writer, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener las credenciales de inicio de sesión desde el cuerpo de la solicitud
	var creds Credentials
	err := json.NewDecoder(request.Body).Decode(&creds)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//verificar las credenciales
	db := commons.GetConnection()
	db.Where("user_name = ? AND password = ?", creds.UserName, creds.Password).Find(&user)

	if user.ID != uuid.Nil {
		//genero token
		token := "token de autenticacion"

		//envio token en la respuesta
		json.NewEncoder(writer).Encode(Token{Token: token})
		// commons.SendResponse(writer, http.StatusOK, []byte(token))
	} else {
		commons.SendError(writer, http.StatusNoContent)
	}
}
