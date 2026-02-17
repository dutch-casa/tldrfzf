package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	list, err := exec.Command("tldr", "--list").Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tldr --list: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("fzf",
		"--ansi",
		"--preview", `tldr {} | bat --plain --color=always --language=markdown`,
		"--preview-window", "right:70%:wrap",
		"--prompt", "tldr> ",
	)
	cmd.Stdin = strings.NewReader(string(list))
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		var exit *exec.ExitError
		if errors.As(err, &exit) {
			// fzf: 1 = no match, 130 = interrupted by user
			if exit.ExitCode() == 1 || exit.ExitCode() == 130 {
				os.Exit(0)
			}
		}
		fmt.Fprintf(os.Stderr, "fzf: %v\n", err)
		os.Exit(1)
	}

	selected := strings.TrimSpace(string(out))
	if selected == "" {
		return
	}

	show := exec.Command("tldr", selected)
	show.Stdout = os.Stdout
	show.Stderr = os.Stderr
	if err := show.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "tldr %s: %v\n", selected, err)
		os.Exit(1)
	}
}
