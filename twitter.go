package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func getTwitterAPI() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}

func postTweet(text string) error {
	api := getTwitterAPI()
	_, err := api.PostTweet(text, nil)
	return err
}
