package vim

import (
	"os"
	"os/exec"
	"slices"
	"strings"
)

func ServersList() ([]string, error) {
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

func NewServer(serverName string, vimFlags []string) (string, error) {
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

func OpenServer(serverName string, vimFlags []string) (string, error) {
	newVimArgs := append([]string{"--remote-tab-silent", serverName}, vimFlags...)

	return NewServer(serverName, newVimArgs)
}
