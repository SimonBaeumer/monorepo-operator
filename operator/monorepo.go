package operator

import (
    "fmt"
    "github.com/SimonBaeumer/cmd"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "path"
)

type MonoRepo struct {
    Projects     []Project `yaml:"projects"`
    OperatingDir string    `yaml:"operating-directory"`
}

func NewMonoRepo(config string) (*MonoRepo, error) {
    m := &MonoRepo{}

    out, _ := ioutil.ReadFile(config)
    err := yaml.Unmarshal(out, m)
    if err != nil {
        return &MonoRepo{}, err
    }

    return m, nil
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

func (m *MonoRepo) ExecOnProject(p Project, command string) {
    fmt.Println("> Execute on " + p.Name)

    setWorkingDir := func (c *cmd.Command) {
        c.WorkingDir = path.Join(m.OperatingDir, p.Name)
    }

    c := cmd.NewCommand(
        command,
        cmd.WithStandardStreams,
        setWorkingDir,
    )

    err := c.Execute()
    if err != nil {
        panic(err.Error())
    }
}

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