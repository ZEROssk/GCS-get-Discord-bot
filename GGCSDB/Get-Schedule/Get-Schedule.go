package GetSchedule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func readClientJSON(file string) (*oauth2.Config) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(f, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config
}

func getClient(config *oauth2.Config, tokFile string) *http.Client {
	f, err := os.Open(tokFile)
	if err != nil {
		fmt.Printf("Open tokFile error: %v", err)
	}
	defer f.Close()
	tok := &oauth2.Token{}
	json.NewDecoder(f).Decode(tok)
	return config.Client(context.Background(), tok)
}

func getEvents(sv *calendar.Service, date string, min string, max string) *calendar.Events {
	Ev, err := sv.Events.List("primary").TimeMin(date + min).TimeMax(date + max).
	SingleEvents(true).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("func getEvents erroe: %v", err)
	}
	return Ev
}

func checkHoliday() string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(location)
	newdate := t.AddDate(0, 0, 0)
	weekday := newdate.Weekday()
	check := weekday.String()
	return check
}

func Get_Sc_Today(s *discordgo.Session, m *discordgo.MessageCreate) string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(location).Format(time.RFC3339)
	today_date := t[:11]
	Date := t[:10]

	min_time := "1:00:00+09:00"
	max_time := "23:00:00+09:00"

	secretJSON := "./TokenFile/secret.json"
	clientJSON := "./TokenFile/credentials.json"

	config := readClientJSON(clientJSON)
	client := getClient(config, secretJSON)

	h := "\n"
	var Sc string

	checkHoli := checkHoliday()
	if checkHoli == "Saturday" || checkHoli == "Sunday" {
		Sc = "Today is " + checkHoli

		embed := &discordgo.MessageEmbed{
			Title:	"Today Schedule",
			Color:	0xd3381c,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:	Date,
					Value:	Sc,
					Inline: false,
				},
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		return checkHoli
	}

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	events := getEvents(srv, today_date, min_time, max_time)

	if len(events.Items) == 0 {
		Sc = "No schedule"
	} else {
		for _, item := range events.Items {
			d := item.Start.DateTime
			SdateTime := ""
			if d == "" {
				SdateTime = "00:00"
			} else {
				SdateTime = d[11:16]
			}

			Sc += SdateTime + " " + item.Summary + h
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:	"Today Schedule",
		Color:	0x00cc66,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:	Date,
				Value:	Sc,
				Inline: false,
			},
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
	return "No schedule"
}

func Get_Sc_Week(s *discordgo.Session, m *discordgo.MessageCreate) string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(location)

	min_time := "1:00:00+09:00"
	max_time := "23:00:00+09:00"

	secretJSON := "./TokenFile/secret.json"
	clientJSON := "./TokenFile/credentials.json"

	config := readClientJSON(clientJSON)
	client := getClient(config, secretJSON)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	var FieldsSlice []*discordgo.MessageEmbedField
	h := "\n"

	for i := 0; i < 5; i++ {
		newdate := t.AddDate(0, 0, i)
		weekday := newdate.Weekday()
		check := weekday.String()

		fdate := newdate.Format(time.RFC3339)
		date := fdate[:11]
		Date := fdate[:10]

		events := getEvents(srv, date, min_time, max_time)
		var Sc string

		checkHoli := checkHoliday()
		if checkHoli == "Saturday" || checkHoli == "Sunday" {
			Sc = "Today is " + checkHoli
	
			embed := &discordgo.MessageEmbed{
				Title:	"Week Schedule",
				Color:	0xd3381c,
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:	Date,
						Value:	Sc,
						Inline: false,
					},
				},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
			return checkHoli
		}

		if len(events.Items) == 0 {
			Sc = "No schedule"
		} else {
			for _, item := range events.Items {
				d := item.Start.DateTime
				SdateTime := ""
				if d == "" {
					SdateTime = "00:00"
				} else {
					SdateTime = d[11:16]
				}

				Sc += SdateTime + " " + item.Summary + h
			}
		}

		em := &discordgo.MessageEmbedField{
			Name: Date,
			Value: Sc,
			Inline: false,
		}
		FieldsSlice = append(FieldsSlice, em)

		if check == "Friday" {
			embed := &discordgo.MessageEmbed{
				Title:	"Week Schedule",
				Color:	0x00cc66,
				Fields: FieldsSlice,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
			break
		}
	}
	return "No schedule"
}

func Get_Sc_NWeek(s *discordgo.Session, m *discordgo.MessageCreate) string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(location)

	min_time := "1:00:00+09:00"
	max_time := "23:00:00+09:00"

	secretJSON := "./TokenFile/secret.json"
	clientJSON := "./TokenFile/credentials.json"

	config := readClientJSON(clientJSON)
	client := getClient(config, secretJSON)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	var NMdate time.Time
	var FieldsSlice []*discordgo.MessageEmbedField
	h := "\n"

	for i := 0; i < 7; i++ {
		nextMonD := t.AddDate(0, 0, i+1)
		day := nextMonD.Weekday()
		checkWDay := day.String()
		if checkWDay == "Monday" {
			NMdate = nextMonD
			break
		}
	}

	for i := 0; i < 5; i++ {
		newdate := NMdate.AddDate(0, 0, i)
		weekday := newdate.Weekday()
		check := weekday.String()

		fdate := newdate.Format(time.RFC3339)
		date := fdate[:11]
		Date := fdate[:10]

		events := getEvents(srv, date, min_time, max_time)
		var Sc string

		if len(events.Items) == 0 {
			Sc = "No schedule"
		} else {
			for _, item := range events.Items {
				d := item.Start.DateTime
				SdateTime := ""
				if d == "" {
					SdateTime = "00:00"
				} else {
					SdateTime = d[11:16]
				}

				Sc += SdateTime + " " + item.Summary + h
			}
		}

		em := &discordgo.MessageEmbedField{
			Name: Date,
			Value: Sc,
			Inline: false,
		}
		FieldsSlice = append(FieldsSlice, em)

		if check == "Friday" {
			embed := &discordgo.MessageEmbed{
				Title:	"Next Week Schedule",
				Color:	0x00cc66,
				Fields: FieldsSlice,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
			break
		}
	}
	return "No schedule"
}

