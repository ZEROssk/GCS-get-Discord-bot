package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"os/signal"
	"syscall"
	"reflect"

	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Authentication"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/SendMessage"
	"github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/SendMessageRegular"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Regular_execution(bot *discordgo.Session) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("ERROR: Failed to LoadLocation:", err)
	}
	Reg := func() {
		bot.AddHandler(SendMessageRegular.SendMRegular)
		err = bot.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
			return
		}
	}

	diff := 3 * time.Second

	ticker := time.NewTicker(diff)

	for {
		select {
		case <-ticker.C:
			Reg()
		}
	}
}

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

	go Regular_execution(bot)

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

