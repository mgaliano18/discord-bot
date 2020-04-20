package main

import (
	"fmt"
	"github.com/mgaliano18/discord-bot/bot"
	"github.com/mgaliano18/discord-bot/config"
)

const Version = "0.0.1"

func main(){
	fmt.Println(fmt.Sprintf("Bot start - Version %s", Version))

	config, err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start(config)

	<-make(chan struct{})
	return
}
