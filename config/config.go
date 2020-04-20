package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SpotifyConfig struct {
	RedirectURL string `json:"RedirectURL"`
	ClientID    string `json:"ClientID"`
	SecretKey   string `json:"SecretKey"`
	State       string `json:"State"`
}

type DiscordConfig struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	BotID     string `json:"BotID"`
	Folder    string `json:"Folder"`
}

func ReadDiscordConfig() (*DiscordConfig, error) {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./discord.config.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(file))
	var config *DiscordConfig
	config = &DiscordConfig{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return config, nil
}