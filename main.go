package main

import (
	"github.com/elcompadre/elcompadre/copro/config"
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/lib/pq"
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
        "host="+ conf.Postgre.Host + " port=" + conf.Postgre.Port +" user="+ conf.Postgre.User +
        " dbname="+ conf.Postgre.Dbname +" sslmode=disable password="+ 
        conf.Postgre.Password)

    if err != nil {
        panic("failed to connect database")
    } else {
        fmt.Println("Database connected")
    }

    defer db.Close()

    db.AutoMigrate(&Resource{})

    router.HandleFunc("/resources", GetResources).Methods("GET")
    // router.HandleFunc("/resources/{id}", GetResource).Methods("GET")
    // router.HandleFunc("/resources", CreateResource).Methods("POST")
    // router.HandleFunc("/resources/{id}", DeleteResource).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":"+ conf.Postgre.SiteHost , router))
}

func GetResources(w http.ResponseWriter, r *http.Request) {
    var resources []Resource
    db.Find(&resources)
    json.NewEncoder(w).Encode(&resources)
}
