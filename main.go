package main

import (
	"time"
	"strconv"
	"github-notifier/helpers"
	"os"
	"os/exec"
	"os/user"
	"io/ioutil"
	"syscall"
)

const pidFileName = ".github-notifier.pid"

func getHomeDir() (string, error) {
	usr, err := user.Current()

	return usr.HomeDir, err
}

func start() {
	// Get home dir
	homeDir, usrErr := getHomeDir()
	if usrErr != nil {
		panic(usrErr)
	}

	// Check if app doesn't already running
	if _, err := os.Stat(homeDir + "/" + pidFileName); err == nil {
		panic("GitHub notifier already running")
	}

	// Run this app again
	cmd := exec.Command(os.Args[0], os.Args[2:]...)

	curDir, osErr := os.Getwd()

	if osErr != nil {
		panic(osErr)
	}

	cmd.Dir = curDir
	cmd.Start()

	// Store PID in file
	pidFileErr := ioutil.WriteFile(homeDir+"/"+pidFileName, []byte(strconv.Itoa(cmd.Process.Pid)), os.FileMode(0644))

	if pidFileErr != nil {
		cmd.Process.Kill()
		panic(pidFileErr)
	}

	cmd.Process.Release()
}

func stop() {
	// Read PID file
	homeDir, usrErr := getHomeDir()
	if usrErr != nil {
		panic(usrErr)
	}

	content, err := ioutil.ReadFile(homeDir + "/" + pidFileName)

	if err != nil {
		panic(err)
	}

	pid, err := strconv.Atoi(string(content))

	if err != nil {
		panic(err)
	}

	// Kill with PID from PID file
	syscall.Kill(pid, 15)

	// Remove PID file
	removeErr := os.Remove(homeDir + "/.github-notifier.pid")

	if removeErr != nil {
		panic(removeErr)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "start" {
		start()
	} else if len(os.Args) > 1 && os.Args[1] == "stop" {
		stop()
	} else {
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
