package operator

import (
	"github.com/SimonBaeumer/cmd"
	"log"
	"strings"
)

// RemoteBranches returns a list of all branches from the given git repository path
func RemoteBranches(workingDir string) []string {
	c := createCommand("git branch -r | grep -v 'HEAD -> origin' | awk -F '/' '{print $2}'", workingDir)

	return strings.Split(c.Stdout(), "\n")
}

// LocalBranches returns a list of local branches
func LocalBranches(workingDir string) []string {
	command := "git for-each-ref --format='%(refname:short)' refs/heads/*"
	c := createCommand(command, workingDir)

	return strings.Split(c.Stdout(), "\n")

}

func createCommand(command string, workingDir string) *cmd.Command {
	setWorkingDir := func(c *cmd.Command) {
		c.WorkingDir = workingDir
	}
	c := newCommand(command, setWorkingDir)
	err := c.Execute()
	if err != nil {
		log.Fatal(err)
	}
	return c
}
