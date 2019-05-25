package SendMessage

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Get-Schedule"
)

var (
	get = "!get"
	set = "!set"
	man = "!man"
	manM = "```
			!set 指定した時間に予定を投稿するチャンネルを設定します
			!get 予定を取得して投稿します

			!man ヘルプです
			```"
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
		schedule := GetSchedule.Get_Sc()
		s.ChannelMessageSend(m.ChannelID, "チャンネルを設定しました")
		SetChannel(m.ChannelID)

	case m.Content == man:
		s.ChannelMessageSend(m.ChannelID, manM)
	}
}

func SetChannel(id *discordgo.Channel) {
	file, err := os.OpenFile("ID.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, id)
}

