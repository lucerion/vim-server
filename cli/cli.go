package cli

import (
	"flag"
	"fmt"
	"os"
	"github.com/lucerion/vim-server/vim"
)

type Config struct {
	Auto bool
}

func ParseFlags() (Config, []string) {
	auto := flag.Bool("auto", false, "connect automatically if only one server runned")
	help := flag.Bool("help", false, "show help message")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	return Config{Auto: *auto}, flag.Args()
}

func Ask(message string) string {
	var input string

	fmt.Println(message)
	fmt.Scan(&input)

	return input
}

func StartServer(serverName string, vimArgs []string) {
	_, err := vim.StartServer(serverName, vimArgs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func OpenServer(serverName string, vimArgs []string) {
	_, err := vim.OpenServer(serverName, vimArgs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func SelectServer(serversList []string, vimArgs []string) {
	printServers(serversList)
	serverName := Ask("Enter server name:")
	if vim.IsServerExists(serverName) {
		OpenServer(serverName, vimArgs)
	} else {
		StartServer(serverName, vimArgs)
	}
}

func printServers(serversList []string) {
	fmt.Println("Servers:")
	for _, serverName := range serversList {
		fmt.Println(serverName)
	}
}
