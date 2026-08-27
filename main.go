package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"github.com/lucerion/vim-server/cli"
	"github.com/lucerion/vim-server/vim"
)

func main() {
	if _, err := exec.LookPath("vim"); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	config, vimFlags := cli.ParseFlags()

	serversList, err := vim.ServersList()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(serversList) == 0 {
		serverName := cli.Ask("Enter new server name:")
		newServer(serverName, vimFlags)
	}

	if len(serversList) == 1 && config.Auto {
		serverName := serversList[0]
		openServer(serverName, vimFlags)
	}

	printServers(serversList)
	serverName := cli.Ask("Enter new or existing server name:")
	selectServer(serverName, serversList, vimFlags)
}

func newServer(serverName string, vimFlags []string) {
	_, err := vim.NewServer(serverName, vimFlags)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func openServer(serverName string, vimFlags []string) {
	_, err := vim.OpenServer(serverName, vimFlags)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func selectServer(serverName string, serversList []string, vimFlags []string) {
	isServerExists := slices.Contains(serversList, strings.ToUpper(serverName))

	if isServerExists  {
		openServer(serverName, vimFlags)
	} else {
		newServer(serverName, vimFlags)
	}
}

func printServers(serversList []string) {
	fmt.Println("Servers:")

	for _, serverName := range serversList {
		fmt.Println(serverName)
	}
}
