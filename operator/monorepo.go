package operator

import (
    "fmt"
    "github.com/SimonBaeumer/cmd"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
    "path"
)

type MonoRepo struct {
    Projects     []Project `yaml:"projects"`
    OperatingDir string    `yaml:"operating-directory"`
}

// NewMonoRepo creates a new instance with the content from the given config file
func NewMonoRepo(config string) (*MonoRepo, error) {
    m := &MonoRepo{}

    out, _ := ioutil.ReadFile(config)
    err := yaml.Unmarshal(out, m)
    if err != nil {
        return &MonoRepo{}, err
    }

    return m, nil
}

// NewMonoRepoFromPath will initialize all directories under a specified path
// as a sub-tree repository
func NewMonoRepoFromPath(gitBaseUrl string, subtreeParentDirectory string, operatingDir string) *MonoRepo {
    monoRepo := MonoRepo{
        OperatingDir: operatingDir,
    }

    projects, _ := ioutil.ReadDir(subtreeParentDirectory)
    for _, f := range projects {
        if !f.IsDir() {
            continue
        }

        p := Project{
            Name:   f.Name(),
            GitUrl: fmt.Sprintf("%s/%s", gitBaseUrl, f.Name()),
            Path:   path.Join(subtreeParentDirectory, f.Name()),
        }

        monoRepo.Add(p)
    }

    return &monoRepo
}

// Add adds a new project to the mono repo
func (m *MonoRepo) Add(p Project) {
    m.Projects = append(m.Projects, p)
}

//Exec executes a command on all subrepos
func (m *MonoRepo) Exec(command string) {
    for _, p := range m.Projects {
        m.ExecOnProject(p, command)
    }
}

// ExecOnProject executes a command on a single project
func (m *MonoRepo) ExecOnProject(p Project, command string) {
    fmt.Println("> Execute on " + p.Name)

    setWorkingDir := func (c *cmd.Command) {
        c.WorkingDir = path.Join(m.OperatingDir, p.Name)
    }

    c := newCommand(
        command,
        cmd.WithStandardStreams,
        setWorkingDir,
    )
    m.exec(c)
}

// GetProject returns a project by name, if no project was found it will return an error.
func (m *MonoRepo) GetProject(name string) (Project, error) {
    for _, p := range m.Projects {
        if p.Name == name {
            return p, nil
        }
    }
    return Project{}, fmt.Errorf("Project %s not found", name)
}

// WriteConfigFile writes the content of the MonoRepo struct to a given config file
func (m *MonoRepo) WriteConfigFile(configFile string) error {
    out, err := yaml.Marshal(m)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(configFile, out, 0755)
    if err != nil {
        return fmt.Errorf("File %s error: %s", configFile, err.Error())
    }
    return nil
}

// Remove removes the operating directory
func (m *MonoRepo) Remove() error {
    return os.RemoveAll(m.OperatingDir)
}

// Clone will clone all repositories which are configured into the operating directory
func (m *MonoRepo) Clone() error {
    for _, p := range m.Projects {
        fmt.Println("> Cloning " + p.Name)
        if err := p.GitClone(m.OperatingDir); err != nil {
           return fmt.Errorf("error while cloning: %s", err.Error())
        }
    }
    return nil
}

// Sync will create subtrees of all projects and create a branch for it
// after that it will be pushed to the remote destination
func (m *MonoRepo) Sync(branch string, useForce bool) {
    forceFlag := ""
    if useForce {
        forceFlag = "-f"
    }
    for _, p := range m.Projects {
        splitBranch := fmt.Sprintf("%s-%s", p.Name, branch)

        // Split project
        fmt.Printf("> split project %s in branch %s\n", p.Name, splitBranch)
        m.SplitProject(p, splitBranch)

        // Add project remote with its' git-url
        fmt.Printf("> add remote %s\n", p.Name)
        addRemoteCmd := newCommand(
            fmt.Sprintf("git remote add %s %s", p.Name, p.GitUrl),
            cmd.WithStandardStreams)
        m.exec(addRemoteCmd)

        // Push project from the split branch to the configured branch
        fmt.Printf("> push project %s\n", p.Name)
        pushCmd := newCommand(
            fmt.Sprintf("git push %s %s %s:%s", forceFlag, p.Name, splitBranch, branch),
            cmd.WithStandardStreams)
        m.exec(pushCmd)

        // Remove created project remote
        fmt.Printf("> remove remote %s\n", p.Name)
        delCmd := newCommand(fmt.Sprintf("git remote rm %s", p.Name), cmd.WithStandardStreams)
        m.exec(delCmd)

        // Remove split branch
        fmt.Printf("> remove branch %s\n", splitBranch)
        delBranchCmd := newCommand(fmt.Sprintf("git branch -D %s", splitBranch), cmd.WithStandardStreams)
        m.exec(delBranchCmd)

        // Print empty line
        fmt.Println()
    }
}

func (m *MonoRepo) exec(c *cmd.Command) {
    if err := c.Execute(); err != nil {
        log.Fatal(err)
    }
}

// SplitProject splits the project and returns the hash or branch name
// If no branch name is given it will only create a hash for the subtree
func (m *MonoRepo) SplitProject(p Project, branch string) string {
    createSplitCmd := fmt.Sprintf("git subtree split -P %s", p.Path)
    if branch != "" {
        createSplitCmd = fmt.Sprintf("%s -b %s", createSplitCmd, branch)
    }

    c := newCommand(createSplitCmd)
    m.exec(c)

    return c.Stdout()
}

// Wrapper function to add some options which will always be needed
func newCommand(command string, options ...func(*cmd.Command)) *cmd.Command {
    parentEnv := func (c *cmd.Command) {
        c.Env = os.Environ()
    }
    return cmd.NewCommand(command, append(options, parentEnv)...)
}