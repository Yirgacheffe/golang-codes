package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var logger = log.New(os.Stdout, "Env: ", log.LstdFlags)

type Env struct {
	User     string `json:"APP_USER"`
	Password string `json:"APP_PASSWORD"`
	QMgr     string `json:"Q_MGR"`
	QName    string `json:"Q_NAME"`
	Host     string `json:"HOST"`
	Port     string `json:"PORT"`
	Channel  string `json:"CHANNEL"`
	Topic    string `json:"TOPIC_NAME"`
}

type Endpoints struct {
	Points []Env `json:"MQ_ENDPOINTS"`
}

const FULL_STRING = -1

var (
	EnvSettings Env
	Q_EPs       Endpoints
)

func init() {
	jsonFile, err := os.Open("./env.json")
	if err != nil {
		log.Fatal(err)
	}

	logger.Println("Successfully open config file")
	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &Q_EPs)

	// Endpoints in json as an array
	// If there are no elements
	// then EnvSettings will be initialized as empty
	if len(Q_EPs.Points) > 0 {
		EnvSettings = Q_EPs.Points[0]
	}

	environmentOverrides()
}

func environmentOverrides() {
	logger.Println("Lookup for Environment Overrides")
	var s string

	overrides := map[string]*string{
		"APP_USER":     &EnvSettings.User,
		"APP_PASSWROD": &EnvSettings.Password,
		"Q_MGR":        &EnvSettings.QMgr,
		"Q_NAME":       &EnvSettings.QName,
		"HOST":         &EnvSettings.Host,
		"PORT":         &EnvSettings.Port,
		"CHANNEL":      &EnvSettings.Channel,
		"TOPIC_NAME":   &EnvSettings.Topic,
	}

	for f, v := range overrides {
		logger.Printf("Trying to override %s", f)
		s = os.Getenv(f)
		if s != "" {
			*v = s
		}
	}
	// -----------------------------------------------------------
}

func (Env) GetConnection(index int) string {
	var points = Q_EPs.Points
	if index == FULL_STRING {
		var conns []string
		for _, p := range points {
			conns = append(conns, p.Host+"("+p.Port+")")
		}
		return strings.Join(conns[:], ",")
	} else {
		return points[index].Host + "(" + points[index].Port + ")"
	}
}

func (Env) LogSettings() {
	logger.Println("Environment Settings as following")

	logger.Printf("Username      is (%s)\n", EnvSettings.User)
	logger.Printf("Host          is (%s)\n", EnvSettings.Host)
	logger.Printf("Port          is (%s)\n", EnvSettings.Port)
	logger.Printf("Queue Manager is (%s)\n", EnvSettings.QMgr)
	logger.Printf("Queue Name    is (%s)\n", EnvSettings.QName)
	logger.Printf("Channel       is (%s)\n", EnvSettings.Channel)
	logger.Printf("Topic         is (%s)\n", EnvSettings.Topic)
	logger.Printf("Connection    is (%s)\n", EnvSettings.GetConnection(FULL_STRING))
}
