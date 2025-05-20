package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/lucerion/vim-server/cli"
	"github.com/lucerion/vim-server/vim"
)

func main() {
	if _, err := exec.LookPath("vim"); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	config, vimArgs := cli.ParseFlags()

	serversList, err := vim.ServersList()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(serversList) == 0 {
		serverName := cli.Ask("Enter server name:")
		cli.StartServer(serverName, vimArgs)
	}

	if len(serversList) == 1 && config.Auto {
		serverName := serversList[0]
		cli.OpenServer(serverName, vimArgs)
	}

	cli.SelectServer(serversList, vimArgs)
}
