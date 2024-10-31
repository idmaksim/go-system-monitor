package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CenterText(text string, width int) string {
	padding := (width - len(text)) / 2
	return fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", padding),
		text,
		strings.Repeat(" ", width-padding-len(text)))
}
