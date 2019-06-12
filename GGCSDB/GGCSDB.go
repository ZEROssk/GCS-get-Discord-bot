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
	clear = "!clear"
	man = "!man"

	setM = "Set Regular execution"
	clearM = "Clear Regular execution"
	manM = "```＊ CmdList: !today, !week, !nweek, !man\r|\r└─ README: https://github.com/ZEROssk/GCS-get-Discord-bot```"

	min_time = "1:00:00+09:00"
	max_time = "23:00:00+09:00"

	secretJSON = "./TokenFile/secret.json"
	clientJSON = "./TokenFile/credentials.json"
)

func SendM(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case m.Content == today:
		GetSchedule.Get_Sc(s, m)

	case m.Content == week:
		GetSchedule.Get_Sc_Week(s, m)

	case m.Content == nweek:
		GetSchedule.Get_Sc_NWeek(s, m)

	case m.Content == man:
		s.ChannelMessageSend(m.ChannelID, manM)
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

