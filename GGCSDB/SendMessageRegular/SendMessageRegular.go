package SendMessageRegular

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Get-Schedule"
)

func SendMRegular(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	schedule := GetSchedule.Get_Sc()
	s.ChannelMessageSend(m.ChannelID, schedule)
}

func ReadID() (id *discordgo.Channel, err error) {
	file, err := os.Open("ID.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	id, err := ioutil.ReadAll(file)
	if err !- nil {
		return err
	}

	return id
}

