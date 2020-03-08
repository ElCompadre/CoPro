package db

import (
	"github.com/elcompadre/elcompadre/copro/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/elcompadre/elcompadre/copro/models"
)

//Database struct
type Database struct {
	PostgresDb  *gorm.DB
}

var db *gorm.DB

func InitializeDB() {
	fmt.Println("Initialize connection endpoint hit")
	conf := config.New()

	con, err := gorm.Open(
	"postgres",
	"host="+conf.Postgre.Host+" user="+conf.Postgre.User+" dbname="+conf.Postgre.Dbname+" sslmode=disable password="+conf.Postgre.Password)

	fmt.Println("Database connecting...")

	if err != nil {
		panic("failed to connect database " + err.Error())
	}
	defer con.Close()
	db = con
	fmt.Println("Database connected")

	db.AutoMigrate(
		models.Base{}, 
		models.Adresse{}, 
		models.CoPro{}, 
		models.CodePostal{}, 
		models.CompteBancaire{}, 
		models.Personne{}, 
		models.Lot{}, 
		models.TypeLot{}, 
		models.User{})
}

func GetDatabase() *Database {
	return &Database{
		PostgresDb: db,
	}
}
