package bot

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"io/ioutil"

	//"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const PlayFile = "play_files"

func (b Bot) playFile(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Author)
	fmt.Println(m.ChannelID)

	channels, err := s.GuildChannels(m.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, channel := range channels {
		fmt.Println(fmt.Sprintf("Channel ID %s, %s", channel.ID, channel.Name))

		for _, recipient := range channel.Recipients {
			fmt.Println(recipient.Username)
		}
	}

	dgv, err := s.ChannelVoiceJoin(m.GuildID, "551928078486732811", false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start loop and attempt to play all files in the given folder
	fmt.Println("Reading Folder: ", b.configurations.Folder)
	files, _ := ioutil.ReadDir(b.configurations.Folder)
	for _, f := range files {
		fmt.Println("PlayAudioFile:", f.Name())
		s.UpdateStatus(0, f.Name())
		path := fmt.Sprintf("%s/%s", b.configurations.Folder, f.Name())
		fmt.Println(path)
		dgvoice.PlayAudioFile(dgv, path, make(chan bool))
	}

	// Close connections
	dgv.Close()
}
