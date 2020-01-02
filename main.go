package main

import (
	"github.com/elcompadre/copro/config"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/lib/pq"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


var db *gorm.DB
var err error

type Resource struct {
    gorm.Model

    Link        string
    Name        string
    Author      string
    Description string
    Tags        pq.StringArray `gorm:"type:varchar(64)[]"`
}

func init() {
    err := godotenv.Load() 
    if err != nil{
		log.Print("No .env file found")
	}
}

func main() {
    router := mux.NewRouter()
    conf := config.New()

    db, err = gorm.Open(
        "postgres",
        "host="+os.Getenv("")+ " port=" + os.Getenv("") +" user="+os.Getenv("")+
        " dbname="+os.Getenv("")+" sslmode=disable password="+ 
        os.Getenv(""))

    if err != nil {
        panic("failed to connect database")
    }

    defer db.Close()

    db.AutoMigrate(&Resource{})

    router.HandleFunc("/resources", GetResources).Methods("GET")
    // router.HandleFunc("/resources/{id}", GetResource).Methods("GET")
    // router.HandleFunc("/resources", CreateResource).Methods("POST")
    // router.HandleFunc("/resources/{id}", DeleteResource).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":"+os.Getenv("SITE_PORT"), router))
}

func GetResources(w http.ResponseWriter, r *http.Request) {
    var resources []Resource
    db.Find(&resources)
    json.NewEncoder(w).Encode(&resources)
}
