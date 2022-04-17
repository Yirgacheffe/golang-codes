package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "Env: ", log.LstdFlags)

type DBEnv struct {
	Host     string `json:"DB_HOST"`
	Port     string `json:"DB_PORT"`
	UserName string `json:"DB_USERNAME"`
	Password string `json:"DB_PASSWORD"`
	DBName   string `json:"DB_NAME"`
}

type DBEndpoints struct {
	Points []DBEnv `json:"DB_ENDPOINTS"`
}

var (
	DBSettings DBEnv
	DBs        DBEndpoints
)

// ---------------------------------------------------------------
func init() {
	jsonFile, err := os.Open("./configs/db.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	logger.Println("Open db config file succeed")
	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &DBs)

	if len(DBs.Points) > 0 {
		DBSettings = DBs.Points[0]
	}

	environmentOverrides()
}

func environmentOverrides() {
	logger.Println("Lookup for Environment Overrides")
	var s string

	overrides := map[string]*string{
		"DB_HOST":     &DBSettings.Host,
		"DB_PORT":     &DBSettings.Port,
		"DB_USERNAME": &DBSettings.UserName,
		"DB_PASSWORD": &DBSettings.Password,
		"DB_NAME":     &DBSettings.DBName,
	}

	for f, v := range overrides {
		logger.Printf("Trying to override %s", f)
		s = os.Getenv(f)
		if s != "" {
			*v = s
		}
	}
}

func (DBEnv) LogSettings() {
	logger.Println("--------- DB Settings as following ---------")

	logger.Printf("Host     is (%s)\n", DBSettings.Host)
	logger.Printf("Port     is (%s)\n", DBSettings.Port)
	logger.Printf("UserName is (%s)\n", DBSettings.UserName)
	logger.Printf("DBName   is (%s)\n", DBSettings.DBName)
}
