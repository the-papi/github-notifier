package main

import (
	"time"
	"strconv"
	"os"
	"go/build"
	"github.com/PapiCZ/github-notifier/helpers"
	"github.com/PapiCZ/github-notifier/settings"
	"os/user"
)


func setHomeDir() {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	settings.HomeDir = usr.HomeDir
}

func main() {
	setHomeDir()

	// Get GOPATH
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}

	if len(os.Args) > 1 {
		cmd := helpers.NewCommand(settings.HomeDir, goPath, settings.SrcRoot)

		switch os.Args[1] {
		case "start":
			cmd.Start(settings.PidFileName)
			break
		case "stop":
			cmd.Stop(settings.PidFileName)
			break
		case "install":
			cmd.Install()
			break
		}
	} else {
		runApp(settings.HomeDir, settings.ConfigFileName)
	}
}

func runApp(homeDir string, configFileName string) {
	config := helpers.NewConfig(homeDir + settings.ConfigPath + "/" + configFileName)

	github := helpers.NewGithubNotifier(config.Get("api_token"))

	interval, err := strconv.ParseInt(config.Get("interval"), 10, 0)

	if err != nil {
		panic(err)
	}

	github.ListenToNotifications(time.Duration(interval) * time.Second)
}
