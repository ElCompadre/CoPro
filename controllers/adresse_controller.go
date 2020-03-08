package controllers

import (
	"encoding/json"
	"github.com/elcompadre/elcompadre/copro/models"
	"github.com/elcompadre/elcompadre/copro/db"
	"github.com/elcompadre/elcompadre/copro/utils"
	"net/http"
	"fmt"
)

func GetAddresses(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Get addresses")
	db := db.GetDatabase()
	var count int
	var responseMessage string
	addresses := db.PostgresDb.Find(models.Adresse{}).Count(&count)
	if count > 0 {
		responseMessage = "Valid request"
	} else {
		responseMessage = "No addresses find"
	}
	utils.Respond(responseWriter, utils.Message(count > 0, responseMessage))
	json.NewEncoder(responseWriter).Encode(addresses)
}

func CreateAddress(responseWriter http.ResponseWriter, request *http.Request) {
	db := db.GetDatabase()
	var newAddress *models.Adresse
	err := json.NewDecoder(request.Body).Decode(&newAddress)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		utils.Respond(responseWriter, utils.Message(false, "Invalid request"))
		return
	}
	db.PostgresDb.Create(&newAddress)
	
	if db.PostgresDb.NewRecord(newAddress) {
		responseWriter.WriteHeader(http.StatusCreated)
	} else {
		responseWriter.WriteHeader(http.StatusBadRequest)
	}
	utils.Respond(responseWriter, utils.Message(true, "Address created"))
	json.NewEncoder(responseWriter).Encode(newAddress)
}

func UpdateAddress(responseWriter http.ResponseWriter, request *http.Request) {
	db := db.GetDatabase()
	var updateAddress *models.Adresse
	err := json.NewDecoder(request.Body).Decode(&updateAddress)
	if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
		utils.Respond(responseWriter, utils.Message(false, "Invalid request"))
		return
	}
	db.PostgresDb.Update(updateAddress)

	utils.Respond(responseWriter, utils.Message(true, "Address updated"))
	json.NewEncoder(responseWriter).Encode(updateAddress)
}