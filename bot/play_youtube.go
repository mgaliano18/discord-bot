package bot

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"io/ioutil"
	//"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const Play = "play"

func (b Bot) play(s *discordgo.Session, m *discordgo.MessageCreate, values ...string) {
	dgv, err := s.ChannelVoiceJoin(m.GuildID, "551928078486732811", false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start loop and attempt to play all files in the given folder
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println("PlayAudioFile:", f.Name())
		s.UpdateStatus(0, f.Name())
		path := fmt.Sprintf("%s/%s", "./", f.Name())
		fmt.Println(path)
		dgvoice.PlayAudioFile(dgv, path, make(chan bool))
	}

	// Close connections
	dgv.Close()
}
