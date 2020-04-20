package main

import (
	"fmt"
	"github.com/mgaliano18/discord-bot/bot"
	"github.com/mgaliano18/discord-bot/config"
)

const Version = "0.0.1"

func main() {
	fmt.Println(fmt.Sprintf("Bot start - Version %s", Version))

	config, err := config.ReadDiscordConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot, err := bot.Start(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goBot.Init()

	<-make(chan struct{})
	return
}
