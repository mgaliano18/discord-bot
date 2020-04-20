package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/mgaliano18/discord-bot/config"
)

var MatiServer = "MatiServer"

type Bot struct {
	configurations *config.DiscordConfig
	goBot *discordgo.Session
}

func (b Bot) Init() {
	b.goBot.AddHandler(b.messageHandler)
}


func Start(config *config.DiscordConfig) (*Bot, error){
	goBot, err := discordgo.New("Bot " + config.Token)
	goBot.UserAgent = MatiServer
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	config.BotID = u.ID
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("Bot is started!")

	return &Bot{
		configurations: config,
		goBot:          goBot,
	}, nil
}
