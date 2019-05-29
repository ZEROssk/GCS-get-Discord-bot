package main

import (
	"fmt"
//	"time"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	fmt.Println("check1")
	bot.AddHandler(SendMessageRegular.SendMRegular)
}

//func Diff_time() {
//	hour := 
//	location, err := time.LoadLocation("Asia/Tokyo")
//	NowTime := time.Now().In(location)
//	noti := time.Date(
//		NowTime.Year(),
//		NowTime.Month(),
//		NowTime.Day(),
//		hour,
//		0,
//		0,
//		0,
//		location,
//	)
//	if noti.Before(NowTime) {
//		noti = noti.Add(time.Hour * 24)
//	}
//
//	diff := noti.Sub(NowTime)
//	return diff
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

