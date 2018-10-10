package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
)

const (
	//twitter API constants
	TWITTERTOKENENV = "TWITTER_ACCESS_TOKEN"
	//openweathermap API constants
	OMWKEY = "OWM_API_KEY"
)

func main() {
	twitterAccessToken := os.Getenv(TWITTERTOKENENV)

	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	client.DoStuff()
}
