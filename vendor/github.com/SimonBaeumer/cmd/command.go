package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

//Command represents a single command which can be executed
type Command struct {
	Command      string
	Env          []string
	Dir          string
	Timeout      time.Duration
	StderrWriter io.Writer
	StdoutWriter io.Writer
	WorkingDir   string
	executed     bool
	exitCode     int
	// stderr and stdout retrieve the output after the command was executed
	stderr bytes.Buffer
	stdout bytes.Buffer
}

// NewCommand creates a new command
// You can add option with variadic option argument
//
// Example:
//      c := cmd.NewCommand("echo hello", function (c *Command) {
//		    c.WorkingDir = "/tmp"
//      })
//      c.Execute()
//
// or you can use existing options functions
//
//      c := cmd.NewCommand("echo hello", cmd.WithStandardStreams)
//      c.Execute()
//
func NewCommand(cmd string, options ...func(*Command)) *Command {
	c := &Command{
		Command:  cmd,
		Timeout:  1 * time.Minute,
		executed: false,
		Env:      []string{},
	}

	c.StdoutWriter = &c.stdout
	c.StderrWriter = &c.stderr

	for _, o := range options {
		o(c)
	}

	return c
}

// WithStandardStreams creates a command which steams its output to the stderr and stdout streams of the operating system
func WithStandardStreams(c *Command) {
	c.StdoutWriter = os.Stdout
	c.StderrWriter = os.Stderr
}

// AddEnv adds an environment variable to the command
// If a variable gets passed like ${VAR_NAME} the env variable will be read out by the current shell
func (c *Command) AddEnv(key string, value string) {
	vars := parseEnvVariableFromShell(value)
	for _, v := range vars {
		value = strings.Replace(value, v, os.Getenv(removeEnvVarSyntax(v)), -1)
	}

	c.Env = append(c.Env, fmt.Sprintf("%s=%s", key, value))
}

// Removes the ${...} characters
func removeEnvVarSyntax(v string) string {
	return v[2:(len(v) - 1)]
}

//Read all environment variables from the given value
//with the syntax ${VAR_NAME}
func parseEnvVariableFromShell(val string) []string {
	reg := regexp.MustCompile(`\$\{.*?\}`)
	matches := reg.FindAllString(val, -1)
	return matches
}

//SetTimeoutMS sets the timeout in milliseconds
func (c *Command) SetTimeoutMS(ms int) {
	if ms == 0 {
		c.Timeout = 1 * time.Minute
		return
	}
	c.Timeout = time.Duration(ms) * time.Millisecond
}

// SetTimeout sets the timeout given a time unit
// Example: SetTimeout("100s") sets the timeout to 100 seconds
func (c *Command) SetTimeout(timeout string) error {
	d, err := time.ParseDuration(timeout)
	if err != nil {
		return err
	}

	c.Timeout = d
	return nil
}

//Stdout returns the output to stdout
func (c *Command) Stdout() string {
	c.isExecuted("Stdout")
	return c.stdout.String()
}

//Stderr returns the output to stderr
func (c *Command) Stderr() string {
	c.isExecuted("Stderr")
	return c.stderr.String()
}

//ExitCode returns the exit code of the command
func (c *Command) ExitCode() int {
	c.isExecuted("ExitCode")
	return c.exitCode
}

//Executed returns if the command was already executed
func (c *Command) Executed() bool {
	return c.executed
}

func (c *Command) isExecuted(property string) {
	if !c.executed {
		panic("Can not read " + property + " if command was not executed.")
	}
}

// Execute executes the command and writes the results into it's own instance
// The results can be received with the Stdout(), Stderr() and ExitCode() methods
func (c *Command) Execute() error {
	cmd := createBaseCommand(c)
	cmd.Env = c.Env
	cmd.Dir = c.Dir
	cmd.Stdout = c.StdoutWriter
	cmd.Stderr = c.StderrWriter
	cmd.Dir = c.WorkingDir

	err := cmd.Start()
	if err != nil {
		return err
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			c.getExitCode(err)
			break
		}
		c.exitCode = 0
	case <-time.After(c.Timeout):
		if err := cmd.Process.Kill(); err != nil {
			return fmt.Errorf("Timeout occurred and can not kill process with pid %v", cmd.Process.Pid)
		}
		return fmt.Errorf("Command timed out after %v", c.Timeout)
	}

	//Remove leading and trailing whitespaces
	c.executed = true

	return nil
}

func (c *Command) getExitCode(err error) {
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			c.exitCode = status.ExitStatus()
		}
	}
}
