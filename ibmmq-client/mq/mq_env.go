package mq

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var logger = log.New(os.Stdout, "Env: ", log.LstdFlags)

type MQEnv struct {
	User     string `json:"APP_USER"`
	Password string `json:"APP_PASSWORD"`
	QMgr     string `json:"Q_MGR"`
	QName    string `json:"Q_NAME"`
	Host     string `json:"HOST"`
	Port     string `json:"PORT"`
	Channel  string `json:"CHANNEL"`
	Cipher   string `json:"CIPHER"`
	KeyRepo  string `json:"KEY_REPO"`
}

type MQEndpoints struct {
	Points []MQEnv `json:"MQ_ENDPOINTS"`
}

const FULL_STRING = -1

var (
	MQSettings MQEnv
	Q_EPs      MQEndpoints
)

// ---------------------------------------------------------------
func init() {
	jsonFile, err := os.Open("./envs/mq.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	logger.Println("Open mq config file succeed")
	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &Q_EPs)

	// Endpoints in json as an array
	// If there are no elements
	// then EnvSettings will be initialized as empty
	if len(Q_EPs.Points) > 0 {
		MQSettings = Q_EPs.Points[0]
	}

	environmentOverrides()
}

func environmentOverrides() {
	logger.Println("Lookup for Environment Overrides")
	var s string

	overrides := map[string]*string{
		"APP_USER":     &MQSettings.User,
		"APP_PASSWROD": &MQSettings.Password,
		"Q_MGR":        &MQSettings.QMgr,
		"Q_NAME":       &MQSettings.QName,
		"HOST":         &MQSettings.Host,
		"PORT":         &MQSettings.Port,
		"CHANNEL":      &MQSettings.Channel,
		"CIPHER":       &MQSettings.Cipher,
		"KEY_REPO":     &MQSettings.KeyRepo,
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

func (MQEnv) GetConnection(index int) string {
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

func (MQEnv) LogSettings() {
	logger.Println("--------- MQ Settings as following ---------")

	logger.Printf("Username      is (%s)\n", MQSettings.User)
	logger.Printf("Queue Manager is (%s)\n", MQSettings.QMgr)
	logger.Printf("Queue Name    is (%s)\n", MQSettings.QName)

	logger.Printf("Host          is (%s)\n", MQSettings.Host)
	logger.Printf("Port          is (%s)\n", MQSettings.Port)
	logger.Printf("Connection    is (%s)\n", MQSettings.GetConnection(FULL_STRING))

	logger.Printf("Channel       is (%s)\n", MQSettings.Channel)
	logger.Printf("Cipher        is (%s)\n", MQSettings.Cipher)
	logger.Printf("Key Repo      is (%s)\n", MQSettings.KeyRepo)
}
