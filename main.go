package main

import (
	"time"
	"strconv"
	"os"
	"go/build"
	"github.com/PapiCZ/github-notifier/helpers"
)

const (
	srcRoot = "/src/github.com/PapiCZ/github-notifier/" // Relative to GOPATH
	pidFileName = ".github-notifier.pid"
	configPath  = "./config.json"
)

func main() {
	// Get GOPATH
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}

	if len(os.Args) > 1 {
		cmd := helpers.NewCommand(goPath, srcRoot)

		switch os.Args[1] {
		case "start":
			cmd.Start(pidFileName)
			break
		case "stop":
			cmd.Stop(pidFileName)
			break
		case "install":
			cmd.Install()
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
