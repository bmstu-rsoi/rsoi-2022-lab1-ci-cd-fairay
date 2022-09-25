package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type DBConfiguration struct {
	Type string `json:"type"`
	Name string `json:"name"`

	User     string `json:"user"`
	Password string `json:"password"`

	Host string `json:"host"`
	Port string `json:"port"`
}

type Configuration struct {
	DB      DBConfiguration `json:"db"`
	LogFile string          `json:"log_file"`
	Port    uint16          `json:"port"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig(config_path string) {
	file, err := os.Open(config_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}

	Config.LogFile = fmt.Sprintf(Config.LogFile, Config.Port)
}
