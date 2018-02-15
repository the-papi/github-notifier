package main

import (
	"github.com/google/go-github/github"
	"time"
	"strconv"
	"github-notifier/helpers"
)

var notificationChannel chan *github.Notification

func main() {
	config := helpers.NewConfig("./github.json")

	github := helpers.NewGithubNotifier(config.Get("api_token"))

	interval, err := strconv.ParseInt(config.Get("interval"), 10, 0)

	if err != nil {
		panic(err)
	}

	github.ListenNotifications(&notificationChannel, time.Duration(interval)*time.Second)
}
