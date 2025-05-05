package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/lucerion/vim-server/cli"
	"github.com/lucerion/vim-server/server"
)

func main() {
	if _, err := exec.LookPath("vim"); err != nil {
		fmt.Print("Error: vim is not installed")
		os.Exit(1)
	}

	config, vimArgs := cli.ParseFlags()

	serversList, err := server.List()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if len(serversList) == 0 {
		serverName := cli.Ask("Enter server name:\n")
		startServer(serverName, vimArgs)
	}

	if len(serversList) == 1 && config.Auto {
		serverName := serversList[0]
		openServer(serverName, vimArgs)
	}

	selectServer(serversList, vimArgs)
}

func startServer(serverName string, vimArgs []string) {
	_, err := server.Start(serverName, vimArgs)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func openServer(serverName string, vimArgs []string) {
	_, err := server.Open(serverName, vimArgs)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func selectServer(serversList []string, vimArgs []string) {
	printServers(serversList)
	serverName := cli.Ask("Enter server name:\n")
	if server.Exists(serverName) {
		openServer(serverName, vimArgs)
	} else {
		startServer(serverName, vimArgs)
	}
}

func printServers(serversList []string) {
	fmt.Print("Servers:\n")
	for _, serverName := range serversList {
		fmt.Println(serverName)
	}
}
