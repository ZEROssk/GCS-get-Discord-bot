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

func getEvents(sv *calendar.Service, mindate string, min string, max string) *calendar.Events {
	Ev, err := sv.Events.List("primary").TimeMin(mindate + min).TimeMax(mindate + max).
	SingleEvents(true).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("func getEvents erroe: %v", err)
	}
	return Ev
}

func Get_Sc(s *discordgo.Session, m *discordgo.MessageCreate) string {
	location, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Now().In(location)

	mint := t.Format(time.RFC3339)
	min_date := mint[:11]

	min_time := "3:00:00+09:00"
	max_time := "21:00:00+09:00"

	secretJSON := "./TokenFile/secret.json"
	clientJSON := "./TokenFile/credentials.json"

	config := readClientJSON(clientJSON)
	client := getClient(config, secretJSON)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	for i := 0; i < 5; i++ {
		newdate := t.AddDate(0, 0, i)
		weekday := newdate.Weekday()
		check := weekday.String()

		maxt := newdate.Format(time.RFC3339)
		max_date := maxt[:11]

		events := getEvents(srv, max_date, min_time, max_time)

		if len(events.Items) == 0 {
			s.ChannelMessageSend(m.ChannelID, "No schedule")
		} else {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}

				schedule := item.Summary + " " + date
				fmt.Println(schedule)
				s.ChannelMessageSend(m.ChannelID, schedule)
			}
		}
		if check == "Friday" {
			break
		}
	}
	return "No schedule"
}

