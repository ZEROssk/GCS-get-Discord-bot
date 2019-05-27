package SendMessageRegular

import (
	"os"
	"io/ioutil"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Get-Schedule"
)

func SendMRegular(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println("check")
	Cid := ReadID()
	schedule := GetSchedule.Get_Sc()
	s.ChannelMessageSend(Cid, schedule)
}

func ReadID() string {
	file, err := os.Open("ID.txt")
	if err != nil {
		return ""
	}
	defer file.Close()

	id, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}

	return string(id)
}

