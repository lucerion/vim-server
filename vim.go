package main

import (
	"os"
	"os/exec"
	"slices"
	"strings"
)

func vimServersList() ([]string, error) {
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

func newVimServer(serverName string, vimFlags []string) (string, error) {
	args := []string{"--servername", serverName}

	if len(vimFlags) > 0 {
		args = append(args, vimFlags...)
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

func openVimServer(serverName string, vimFlags []string) (string, error) {
	newVimArgs := append([]string{"--remote-tab-silent", serverName}, vimFlags...)

	return newVimServer(serverName, newVimArgs)
}
