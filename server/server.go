package server

import (
	"os"
	"os/exec"
	"slices"
	"strings"
)

func List() ([]string, error) {
	cmd := exec.Command("vim", "--serverlist")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	str := string(output)
	lines := strings.Split(str, "\n")
	nonEmptyLines := slices.DeleteFunc(lines, func(line string) bool {
		return strings.TrimSpace(line) == ""
	})

	return nonEmptyLines, nil
}

func Start(serverName string, vimArgs []string) (string, error) {
	args := []string{"--servername", serverName}

	if len(vimArgs) > 0 {
		args = append(args, vimArgs...)
	}

	cmd := exec.Command("vim", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return serverName, err
	}

	return serverName, nil
}

func Open(serverName string, vimArgs []string) (string, error) {
	newVimArgs := append([]string{"--remote-tab-silent"}, vimArgs...)

	return Start(serverName, newVimArgs)
}

func Exists(serverName string) bool {
	servers, _ := List()

	return slices.Contains(servers, strings.ToUpper(serverName))
}
