package SendMessage

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Get-Schedule"
)

var (
	get = "!get"
	set = "!set"
	man = "!man"
	manM = "```!get !set !man```"
)

func SendM(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case m.Content == get:
		schedule := GetSchedule.Get_Sc()
		s.ChannelMessageSend(m.ChannelID, schedule)

	case m.Content == set:
		s.ChannelMessageSend(m.ChannelID, "チャンネルを設定しました")

	case m.Content == man:
		s.ChannelMessageSend(m.ChannelID, manM)
	}
}

func SetChannel(id string) {
	file, err := os.OpenFile("ID.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(id)
	file.Write(([]byte)(id))
}

