package helpers

import (
	"io/ioutil"
	"strconv"
	"os"
	"syscall"
	"os/user"
	"os/exec"
)

func getHomeDir() (string, error) {
	usr, err := user.Current()

	return usr.HomeDir, err
}

func Start(pidFileName string) {
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

func Stop(pidFileName string) {
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
