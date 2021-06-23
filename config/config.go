package config

import (
	"os"
	"log"

	"encoding/json"
	"path/filepath"
)

type Configuration struct {
	CONNECTIONSTRING string `json:"url"`
	DATABASE         string `json:"database"`
	User             string `json:"user"`
	Password         string `json:"password"`
}

func GetConfig() Configuration {
	cwd, err := os.Getwd()
	log.Println(err)

	file, err := os.Open(filepath.Join(cwd, "config", "config.json"))
	if err != nil {
		log.Fatal("configuration file not found")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}

	err = decoder.Decode(&config)
	if err != nil {
		log.Println("error:", err)
	}

	return config
}
