package main

import (
	"time"
	"strconv"
	"github-notifier/helpers"
	"os"
)

const (
	pidFileName = ".github-notifier.pid"
	configPath  = "./config.json"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "start" {
		helpers.Start(pidFileName)
	} else if len(os.Args) > 1 && os.Args[1] == "stop" {
		helpers.Stop(pidFileName)
	} else {
		runApp()
	}
}

func runApp() {
	config := helpers.NewConfig(configPath)

	github := helpers.NewGithubNotifier(config.Get("api_token"))

	interval, err := strconv.ParseInt(config.Get("interval"), 10, 0)

	if err != nil {
		panic(err)
	}

	github.ListenToNotifications(time.Duration(interval) * time.Second)
}
