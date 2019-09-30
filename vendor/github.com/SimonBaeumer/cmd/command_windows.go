package cmd

import (
	"os/exec"
	"strings"
)

func createBaseCommand(c *Command) *exec.Cmd {
	cmd := exec.Command(`C:\windows\system32\cmd.exe`, "/C", c.Command)
	return cmd
}

func (c *Command) removeLineBreaks(text string) string {
	return strings.Trim(text, "\r\n")
}
