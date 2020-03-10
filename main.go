package main

import (
	"encoding/json"
	"flag"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	log "github.com/sirupsen/logrus"
)

var (
	logger = log.New()
)

func getTweet(api *anaconda.TwitterApi, tweetID int64) (anaconda.Tweet, error) {
	args := url.Values{}
	args.Add("include_entities", "true")
	args.Add("include_ext_alt_text", "true")
	args.Add("include_card_uri", "true")
	tweet, err := api.GetTweet(tweetID, args)
	if err != nil {
		return anaconda.Tweet{}, err
	}
	return tweet, nil
}

func main() {
	var tweetID = flag.Int64("i", int64(0), "tweet id")
	flag.Parse()

	anaconda.SetConsumerKey(os.Getenv("TC_TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TC_TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TC_TWITTER_ACCESS_TOKEN"), os.Getenv("TC_TWITTER_ACCESS_TOKEN_SECRET"))
	api.SetLogger(anaconda.BasicLogger)

	fmter := new(log.TextFormatter)
	fmter.FullTimestamp = true
	log.SetFormatter(fmter)
	log.SetLevel(log.InfoLevel)

	log.Info("STARTED: ID ", *tweetID)

	tweet, err := getTweet(api, *tweetID)
	if err != nil {
		log.Error("Could not get tweet: ", err)
		os.Exit(-1)
	}

	// print the tweet
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(tweet)
}
