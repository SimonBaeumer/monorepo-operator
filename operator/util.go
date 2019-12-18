package operator

import (
	"github.com/SimonBaeumer/cmd"
	"log"
	"os"
)

func exec(c *cmd.Command) {
	if err := c.Execute(); err != nil {
		log.Fatal(err)
	}

	if c.ExitCode() != 0 {
		log.Fatalf("Received exit code %d, with stderr: \n%s", c.ExitCode(), c.Stderr())
	}
}

// Wrapper function to add some options which will always be needed
func newCommand(command string, options ...func(*cmd.Command)) *cmd.Command {
	parentEnv := func(c *cmd.Command) {
		c.Env = os.Environ()
	}
	return cmd.NewCommand(command, append(options, parentEnv)...)
}
