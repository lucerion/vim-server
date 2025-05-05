package cli

import (
	"flag"
	"fmt"
)

type Config struct {
	Auto bool
}

func ParseFlags() (Config, []string) {
	auto := flag.Bool("auto", false, "connect automatically if only one server runned")
	flag.Parse()

	args := flag.Args()

	return Config{Auto: *auto}, args
}

func Ask(message string) string {
	var input string

	fmt.Print(message)
	fmt.Scan(&input)

	return input
}
