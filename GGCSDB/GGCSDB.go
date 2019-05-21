package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

func main() {
	t := time.Now().Format(time.RFC3339)
	today_date := t[:11]

	min_time := "10:00:00+09:00"
	max_time := "19:00:00+09:00"

	secretJSON := "./Authentication/secret.json"
	clientJSON := "./Authentication/credentials.json"

	config := readClientJSON(clientJSON)
	client := getClient(config, secretJSON)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	events := getEvents(srv, today_date, min_time, max_time)

	if len(events.Items) == 0 {
		fmt.Println("No schedule")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}

