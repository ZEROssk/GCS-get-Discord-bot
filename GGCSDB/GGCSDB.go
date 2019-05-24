package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Authentication"
	"./SendMessage"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

//var (
//	get = "!get"
//	help = "!help"
//)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
//	if m.Author.ID == s.State.User.ID {
//		return
//	}
//
//	switch {
//	case m.Content == get:
//		schedule := GetSchedule.Get_Sc()
//		s.ChannelMessageSend(m.ChannelID, schedule)
//
//	case m.Content == help:
//		s.ChannelMessageSend(m.ChannelID, "HELP")
//	}
//}

func main() {
	Authentication.Auth()
	Env_load()

	TOKEN := "Bot " + os.Getenv("YOUR_TOKEN")
	stop := make(chan os.Signal, 1)

	bot, err := discordgo.New(TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	//bot.AddHandler(messageCreate)
	bot.AddHandler(SendMessage.SendM)
	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.")

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	bot.Close()
}

