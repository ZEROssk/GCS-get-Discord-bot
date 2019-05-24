package SendMessage

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Get-Schedule"
)

var (
	get = "!get"
	help = "!help"
)

func SendM(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case m.Content == get:
		schedule := GetSchedule.Get_Sc()
		s.ChannelMessageSend(m.ChannelID, schedule)

	case m.Content == help:
		s.ChannelMessageSend(m.ChannelID, "HELP")
	}
}

func SendM_Regular(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	schedule := GetSchedule.Get_Sc()
	s.ChannelMessageSend(m.ChannelID, schedule)
}

