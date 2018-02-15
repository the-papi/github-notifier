package main

import (
	"time"
	"strconv"
	"github-notifier/helpers"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] != "--detached" {
		cmd := exec.Command(os.Args[0], "--detached")
		cmd.Dir, _ = os.Getwd()
		cmd.Start()
		cmd.Process.Release()
	} else if os.Args[1] == "--detached" {
		runApp()
	}
}

func runApp() {
	config := helpers.NewConfig("./config.json")

	github := helpers.NewGithubNotifier(config.Get("api_token"))

	interval, err := strconv.ParseInt(config.Get("interval"), 10, 0)

	if err != nil {
		panic(err)
	}

	github.ListenToNotifications(time.Duration(interval) * time.Second)
}
