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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "start":
			helpers.Start(pidFileName)
			break
		case "stop":
			helpers.Stop(pidFileName)
			break
		}
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
