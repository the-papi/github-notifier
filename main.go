package main

import (
	"github.com/google/go-github/github"
	"time"
	"strconv"
	"github-notifier/helpers"
)

func main() {
	config := helpers.NewConfig("./config.json")

	github := helpers.NewGithubNotifier(config.Get("api_token"))

	interval, err := strconv.ParseInt(config.Get("interval"), 10, 0)

	if err != nil {
		panic(err)
	}

	github.ListenToNotifications(time.Duration(interval)*time.Second)
}
