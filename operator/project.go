package operator

import (
	"fmt"
	"github.com/SimonBaeumer/cmd"
	"io/ioutil"
	"os"
	"path"
)

type Project struct {
	Name string `yaml:"name"`
	// Path defines the path to the directory inside the mono repo
	Path string `yaml:"path"`
	// OperatingPath defines the path to the checked out repository, located in the mono-repos` operating dir
	OperatingPath string `yaml:"operating-path,omitempty"`
	GitUrl        string `yaml:"git-url"`
}

// GitClone clones the project into the given destination path
func (p *Project) GitClone(dest string) error {
	cloneCmd := fmt.Sprintf(
		"git clone %s.git %s/%s",
		p.GitUrl,
		dest,
		p.Name,
	)

	exec(newCommand(cloneCmd))

	return nil
}

// Exec executes a command inside the project's directory
func (p *Project) Exec(command string) {
	fmt.Println("> Exec " + p.Name)

	setWorkingDir := func(c *cmd.Command) {
		c.WorkingDir = p.OperatingPath
	}

	c := newCommand(
		command,
		cmd.WithStandardStreams,
		setWorkingDir,
	)

	exec(c)
}

// Lock will lock a project for specific tasks
func (p *Project) Lock() error {
	lockFile := path.Join(p.Path, ".git/monorepo.lock")
	_, e := os.Stat(lockFile)
	if e == nil {
		return fmt.Errorf("repository is already blocked by another operation")
	}

	err := ioutil.WriteFile(lockFile, []byte(``), 0755)
	if err != nil {
		panic(err)
	}
	return nil
}

// Unlock unlocks the repo
func (p *Project) Unlock() {
	lockFile := path.Join(p.Path, ".git/monorepo.lock")
	if err := os.Remove(lockFile); err != nil {
		panic(err)
	}
}
