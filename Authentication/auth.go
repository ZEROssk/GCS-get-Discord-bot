package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"reflect"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func confirmations(name string) error {
	_, err := os.Stat(name)
	fmt.Println(reflect.TypeOf(!os.IsNotExist(err)))
	if !os.IsNotExist(err) {
		fmt.Println("OK")
		return err
	} else {
		fmt.Println("ERROR")
		return err
	}
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code:"+
	"\n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main () {
	client := "credentials.json"
	secret := "secret.json"

	Clierr := confirmations(client)
	if Clierr != nil {
		fmt.Printf("credentials.json undefaind\n")
	} else {
		fmt.Printf("credentials.json done\n")
	}

	Secerr := confirmations(secret)
	if Secerr != nil {
		b, err := ioutil.ReadFile(client)
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		tok := getTokenFromWeb(config)
		saveToken(secret, tok)

	} else {
		fmt.Printf("secret.json is Already exists\n")
	}
}
