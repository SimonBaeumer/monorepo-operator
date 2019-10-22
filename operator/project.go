package operator

import (
	"fmt"
	"github.com/SimonBaeumer/cmd"
	"io/ioutil"
	"os"
	"path"
)

type Project struct {
	Name   string `yaml:"name"`
	Path   string `yaml:"path"`
	GitUrl string `yaml:"git-url"`
}

// GitClone clones the project into the given destination path
func (p *Project) GitClone(dest string) error {
	cloneCmd := fmt.Sprintf(
		"git clone %s.git %s/%s",
		p.GitUrl,
		dest,
		p.Name,
	)

	clone := cmd.NewCommand(cloneCmd)
	err := clone.Execute()
	if err != nil {
		return err
	}

	if clone.ExitCode() != 0 {
		return fmt.Errorf(clone.Stderr())
	}
	return nil
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
