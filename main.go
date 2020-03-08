package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elcompadre/elcompadre/copro/config"
	"github.com/elcompadre/elcompadre/copro/db"
	"github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/elcompadre/elcompadre/copro/controllers"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
	
}

func main() {
	fmt.Println("Copro 2.0")
	router := mux.NewRouter()

	conf := config.New()
	db.InitializeDB()

	router.HandleFunc("/addresses", controllers.GetAddresses).Methods("GET")
	router.HandleFunc("/addresses", controllers.CreateAddress).Methods("POST")
	router.HandleFunc("/addresses/{id}", controllers.UpdateAddress).Methods("PUT")

	fmt.Println(conf.Postgre.SitePort)
	log.Fatal(http.ListenAndServe(":"+ conf.Postgre.SitePort, router))
}

