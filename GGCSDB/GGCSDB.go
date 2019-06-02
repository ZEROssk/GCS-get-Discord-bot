package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"regexp"

	"./Authentication"
	"./Get-Schedule"
//	"github.com/okzk/ticker"
	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

var (
	set = "!set"
	get = "!get"
	clear = "!clear"
	man = "!man"

	setM = "Set Regular execution"
	clearM = "Clear Regular execution"
	manM = "```!get !set !clear !man```"

	Cid string
	check_num int = 0
	Rtime time.Duration = 5
	//kill = make(chan bool)
)

func SendM(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	GetRtime := regexp.MustCompile(`^!set \d\d?`)

	if GetRtime.Match([]byte(m.Content)) {
		splittedCommand := strings.Split(m.Content, " ")

		now := time.Now().In(jst)

		hour, err := strconv.Atoi(splittedCommand[1])
		if err != nil {
			return
		}

		if 24 <= hour {
			return
		}


		notificationTime := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			hour,
			0, // min
			0, // sec
			0, // nsec
			jst,
		)

		// !(now < notification time)
		if notificationTime.Before(now) {
			notificationTime = notificationTime.Add(time.Hour * 24)
		}

		diff := notificationTime.Sub(now)
	}

	switch {
	case m.Content == get:
		schedule := GetSchedule.Get_Sc()
		s.ChannelMessageSend(m.ChannelID, schedule)

	case m.Content == set:
		Cid = m.ChannelID
		check_num = 0
		s.ChannelMessageSend(m.ChannelID, setM)

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

	if Cid == "" && Rtime == 5 {
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
		//ticker := ticker.New(Rtime * time.Second, func(t time.Time) {
		//	s.ChannelMessageSend(Cid, schedule)
		//})
	}
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
	stop := make(chan os.Signal, 1)

	TOKEN := "Bot " + os.Getenv("YOUR_TOKEN")

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

