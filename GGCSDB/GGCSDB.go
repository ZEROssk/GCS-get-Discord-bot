package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"

	"./Authentication"
	"./Get-Schedule"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

var (
	get = "!get"
	help = "!help"
)

//func check_json(name string) error {
//	_, err := os.Stat(name)
//	fmt.Println(reflect.TypeOf(!os.IsNotExist(err)))
//	if !os.IsNotExist(err) {
//		return err}
//	return err
//}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case m.Content == get:
		GetSchedule.Get_Sc()
		s.ChannelMessageSend(m.ChannelID, "DONE")

	case m.Content == help:
		s.ChannelMessageSend(m.ChannelID, "HELP")
	}
}

func main() {
	Authentication.Auth()
//	secret := "./Authentication/secret.json"
//	findJSON := check_json(secret)
//	if findJSON != nil {
//		fmt.Printf("secret.json not found")
//		Auth.auth()
//	}

	Env_load()

	TOKEN := "Bot " + os.Getenv("YOUR_TOKEN")
	stop := make(chan os.Signal, 1)

	bot, err := discordgo.New(TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	bot.AddHandler(messageCreate)
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

