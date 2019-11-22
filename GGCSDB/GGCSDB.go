package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"./Authentication"
	"./Get-Schedule"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

var (
	today = "!today"
	week = "!week"
	nweek = "!nweek"
	man = "!man"
)

func SendM(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:  "Help",
		Color:  0x00cc66,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Today Schedule",
				Value:  "!today",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Week Schedule",
				Value:  "!week",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Next Week Schedule",
				Value:  "!nweek",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Help",
				Value:  "!man",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "README",
				Value:  "`https://github.com/ZEROssk/GCS-get-Discord-bot`",
			},
		},
	}

	switch {
		case m.Content == today:
			GetSchedule.Get_Sc_Today(s, m)
	
		case m.Content == week:
			GetSchedule.Get_Sc_Week(s, m)
	
		case m.Content == nweek:
			GetSchedule.Get_Sc_NWeek(s, m)
	
		case m.Content == man:
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Authentication.Auth()
	Env_load()
	stop := make(chan os.Signal, 1)

	TOKEN := "Bot " + os.Getenv("YOUR_TOKEN")

	bot, err := discordgo.New(TOKEN)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	bot.AddHandler(SendM)

	err = bot.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("---Bot is now running---")

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	bot.Close()
}

