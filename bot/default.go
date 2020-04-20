package bot

import "github.com/bwmarrin/discordgo"

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	switch m.Content {
	case "ping":
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	case PlayFile:
		playFile(s,m)
	}
}