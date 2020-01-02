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
	SiteHost string
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
		 SiteHost: os.Getenv("SITE_PORT"),
	 },
 }
}

// var env string = "prod"

// //GetConfig method
// func (config *Configuration) GetConfig() *Configuration {
// 	if os.Getenv("ENVIRONMENT") == "DEV" {
// 		env = "dev"
// 	}
// 	configFile := fmt.Sprintf("config_%s.json", env)
// 	file, er := os.Open(configFile)
// 	if er != nil {
// 		fmt.Println("error:", er)
// 	}
// 	decoder := json.NewDecoder(file)
// 	defer file.Close()
// 	err := decoder.Decode(&config)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	return config
// }
