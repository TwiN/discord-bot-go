package config

import (
	"encoding/json"
	"os"
	"log"
)

type Configuration struct {
	Users struct {
		OwnerId   string
		BlackList []string
	}
	AutoWelcome bool
}

var Config Configuration

func Load() {
	log.Println("[config][init] Loading configuration file")
	configFile, _ := os.Open("config.json")
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln("[config][init] Unable to decode configuration:", err)
	}
	Config = configuration
}
