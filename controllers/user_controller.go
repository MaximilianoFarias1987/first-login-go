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
