package config

import (

	"os"
)

type PostgreSqlConfig struct {
	Host     string
	Port     string    
	User     string
	Password string 
	Dbname   string 
	SitePort string
}

//Configuration model
type Configuration struct {
	Postgre PostgreSqlConfig
}

func New() *Configuration {
 return &Configuration{
	 Postgre: PostgreSqlConfig{
		 Host: os.Getenv("DB_HOST"),
		 Port: os.Getenv("DB_PORT"),
		 User: os.Getenv("DB_USER"),
		 Password: os.Getenv("DB_PASSWORD"),
		 Dbname: os.Getenv("DBNAME"),
		 SitePort: os.Getenv("SITE_PORT"),
	 },
 }
}
