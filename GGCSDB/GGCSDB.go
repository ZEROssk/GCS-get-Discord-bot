package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"./Authentication"
	"./Get-Schedule"
	"github.com/okzk/ticker"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

var (
	get = "!get"
	set = "!set"
	man = "!man"
	manM = "```!get !set !man```"

	Cid string
	//RegularTime = 7 * time.Second
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
		Cid = m.ChannelID
		s.ChannelMessageSend(m.ChannelID, "チャンネルを設定しました")

	case m.Content == man:
		s.ChannelMessageSend(m.ChannelID, manM)
	}
}

func SendMRegular(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	schedule := GetSchedule.Get_Sc()
	if Cid == "" {
		return
	}

	ticker.New(10 * time.Second, func(t time.Time) {
		s.ChannelMessageSend(Cid, schedule)
	})
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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

	bot.AddHandler(SendM)
	go bot.AddHandler(SendMRegular)

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

