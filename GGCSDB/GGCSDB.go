package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"regexp"
	"strings"
	"strconv"

	"./Authentication"
	"./Get-Schedule"
	//	"github.com/okzk/ticker"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

var (
	get = "!get"
	clear = "!clear"
	man = "!man"

	setM = "Set Regular execution"
	clearM = "Clear Regular execution"
	manM = "```!get, !set <time>, !clear, !man```"

	Cid string
	check_num int = 0
	Rtime time.Duration
)

func SendM(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	GetRtime := regexp.MustCompile(`^!set \d\d?`)

	if GetRtime.Match([]byte(m.Content)) {
		split := strings.Split(m.Content, " ")

		getRt, err := strconv.Atoi(split[1])
		if err != nil {
			return
		}

		Diff_time(getRt)
		Cid = m.ChannelID
		s.ChannelMessageSend(m.ChannelID, setM)
	}

	switch {
	case m.Content == get:
		schedule := GetSchedule.Get_Sc()
		s.ChannelMessageSend(m.ChannelID, schedule)

	case m.Content == clear:
		check_num = 1
		Cid = ""
		Rtime = 0
		s.ChannelMessageSend(m.ChannelID, clearM)
		return

	case m.Content == man:
		s.ChannelMessageSend(m.ChannelID, manM)
	}
}

func SendMRegular(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if Cid == "" {
		return
	}

	schedule := GetSchedule.Get_Sc()

	ticker := time.NewTicker(Rtime * time.Second)

	for {
		select {
		case <-ticker.C:
			if check_num == 1 {
				ticker.Stop()
				return
			}	
			s.ChannelMessageSend(Cid, schedule)
		}
	}
}

func Diff_time(Rt int) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("Error LoadLocation: ", err)
	}

	now := time.Now().In(location)

	if 24 <= Rt {
		return
	}

	notificationTime := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		Rt,
		0, // min
		0, // sec
		0, // nsec
		location,
	)

	if notificationTime.Before(now) {
		notificationTime = notificationTime.Add(time.Hour * 24)
	}

	Rtime = notificationTime.Sub(now)
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
	go bot.AddHandler(SendMRegular)

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

