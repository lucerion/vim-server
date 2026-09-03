package cli

import (
	"fmt"
	"os"
)

const USAGE = `Usage: vim-server [VIM_SERVER_OPTIONS] [OPTIONS] [FILE...]

VIM_SERVER_OPTIONS

  -vs-auto, --vs-auto    connect automatically if only one server runned
  -vs-help, --vs-help    show this help message

OPTIONS & FILES

  all other flags and arguments are passed directly to the vim
`

type Config struct {
	Auto bool
}

func ParseFlags() (Config, []string) {
	var config Config
	var vimFlags []string

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "-vs-auto" || arg == "--vs-auto" {
			config.Auto = true
			continue
		}

		if arg == "-vs-help" || arg == "--vs-help" {
			fmt.Print(USAGE)
			os.Exit(0)
		}

		vimFlags = append(vimFlags, arg)
	}

	return config, vimFlags
}

func Ask(message string) string {
	var input string

	fmt.Println(message)
	fmt.Scan(&input)

	return input
}
