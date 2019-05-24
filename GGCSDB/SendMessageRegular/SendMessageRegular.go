package SendMessageRegular

import (
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

