package bot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (b Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.configurations.BotID {
		return
	}
	if strings.Contains(m.Content, b.configurations.BotPrefix){
		command, values := b.GetCommands(m.Content)
		switch command {
		case "ping":
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		case PlayFile:
			b.playFile(s, m)
		case Play:
			b.play(s,m, values...)
		}
	}
}
