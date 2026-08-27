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

	vimFlags := parseVimFlags()

	return Config{Auto: *auto}, vimFlags
}

func Ask(message string) string {
	var input string

	fmt.Println(message)
	fmt.Scan(&input)

	return input
}

func parseVimFlags() []string {
	var vimFlags []string

	flag.Visit(func(f *flag.Flag) {
		vimFlags = append(vimFlags, fmt.Sprintf("--%s=%s", f.Name, f.Value.String()))
	})

	return vimFlags
}
