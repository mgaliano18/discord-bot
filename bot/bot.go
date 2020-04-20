package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/mgaliano18/discord-bot/config"
)

var config *config.ConfigStruct
var goBot *discordgo.Session
var MatiServer = "MatiServer"

func Start(configToSet config.ConfigStruct) {
	config = configToSet
	goBot, err := discordgo.New("Bot " + config.Token)
	goBot.UserAgent = MatiServer
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	config.BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}