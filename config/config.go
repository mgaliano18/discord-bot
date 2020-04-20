package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Public variables
	Config *ConfigStruct
)

type ConfigStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Folder string `json:"Folder"`
}

func ReadConfig() (*ConfigStruct, error) {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &Config)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return Config, nil
}
