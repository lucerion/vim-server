package cli

import (
	"flag"
	"fmt"
	"os"
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

	fmt.Print(message)
	fmt.Scan(&input)

	return input
}
