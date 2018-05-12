package config

import (
	"encoding/json"
	"os"
	"os/user"
)

type Configuration struct {
	Port       int
	LogPath    string
	SmallFiles bool
	RepoURL    map[string]string
}

func GetConfig() Configuration {
	var configuration Configuration
	usr, err := user.Current()
	var filename string = usr.HomeDir + "/.mongo/config.json"
	//filename is the path to the json config file
	file, err := os.Open(filename)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)

	if err != nil {
		return configuration
	}

	return configuration
}
