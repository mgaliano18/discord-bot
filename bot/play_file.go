package bot

import (
	"fmt"

	//"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const PlayFile = "play_file"

func playFile(s *discordgo.Session, m *discordgo.MessageCreate) {
	dgv, err := s.ChannelVoiceJoin(m.GuildID, "551928078486732811", false,true)
	if err != nil {
		fmt.Println(err)
		return
	}

	//dgvoice.PlayAudioFile(dgv, ,make(chan bool))

	// Close connections
	dgv.Close()
	s.Close()
}
