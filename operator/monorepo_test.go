package operator

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestNewMonoRepo(t *testing.T) {
	f, _ := ioutil.TempFile("/tmp", "tmp_config_monorepo")
	ioutil.WriteFile(f.Name(), []byte(`
projects:
- name: test
  path: tmp/test
  giturl: https://github.com/SimonBaeumer/test
- name: test2
  path: tmp/test2
  giturl: https://github.com/SimonBaeumer/test2
operating-directory: /tmp
`), 0755)
	m, err := NewMonoRepo(f.Name())

	assert.Nil(t, err)
	assert.Len(t, m.Projects, 2)
	assert.Equal(t, m.Projects[0].Name, "test")
	assert.Equal(t, m.OperatingDir, "/tmp")
}

func TestNewMonoRepoFromPath(t *testing.T) {
	operatingDir, _ := ioutil.TempDir("/tmp", "monorepo_test_operating")
	defer os.Remove(operatingDir)
	repoDir, _ := ioutil.TempDir("/tmp", "monorepo_test_repo")
	defer os.Remove(repoDir)

	os.Mkdir(repoDir+"/repo01", 0755)
	os.Mkdir(repoDir+"/repo02", 0755)

	monoRepo := NewMonoRepoFromPath(
		"git@github.com:spf13",
		repoDir,
		operatingDir,
	)

	assert.Equal(t, operatingDir, monoRepo.OperatingDir)
	assert.Equal(t, "repo01", monoRepo.Projects[0].Name)

	assert.Equal(t, "repo02", monoRepo.Projects[1].Name)
	assert.Equal(t, "git@github.com:spf13/repo02.git", monoRepo.Projects[1].GitUrl)
	assert.Equal(t, repoDir+"/repo02", monoRepo.Projects[1].Path)
}

func Test_MonoRepo_GetProject(t *testing.T) {
	m := MonoRepo{
		Projects: []Project{
			{Name: "test"},
			{Name: "test2"},
		},
	}

	p, _ := m.GetProject("test")
	assert.Equal(t, "test", p.Name)

	_, err := m.GetProject("invalid")
	assert.Equal(t, "Project invalid not found", err.Error())
}

func Test_MonoRepo_WriteConfigFile(t *testing.T) {
	f, _ := ioutil.TempFile("/tmp", "monorepo-operator.yml")
	defer os.Remove(f.Name())

	m := MonoRepo{
		Projects: []Project{
			{
				Name:   "project_test",
				GitUrl: "git@github.com:SimonBaeumer/monorepo-operator.git",
			},
		},
		OperatingDir: ".git/subtree",
	}

	err := m.WriteConfigFile(f.Name())

	expected := `projects:
- name: project_test
  path: ""
  git-url: git@github.com:SimonBaeumer/monorepo-operator.git
operating-directory: .git/subtree
`

	assert.Nil(t, err)
	conf, _ := ioutil.ReadFile(f.Name())
	assert.Equal(t, expected, string(conf))
}

func Test_MonoRepo_Exec(t *testing.T) {
	operatingDir, _ := ioutil.TempDir("/tmp", "monorepo_test_operating")
	defer os.Remove(operatingDir)
	repoDir, _ := ioutil.TempDir("/tmp", "monorepo_test_repo")
	defer os.Remove(repoDir)

	os.Mkdir(repoDir+"/repo01", 0755)
	os.Mkdir(repoDir+"/repo02", 0755)
	os.Mkdir(operatingDir+"/repo01", 0755)
	os.Mkdir(operatingDir+"/repo02", 0755)

	m := NewMonoRepoFromPath(
		"git@github.com:SimonBaeumer",
		repoDir,
		operatingDir,
	)

	m.Exec("touch test")

	assert.Len(t, m.Projects, 2)
	for _, p := range m.Projects {
		assert.FileExists(t, path.Join(p.OperatingPath, "test"))
	}
}

func TestMonoRepo_isProtectedBranch(t *testing.T) {
	m := MonoRepo{}
	m.ProtectedBranches = []string{"master", `\d{0,9}\.\p{N}\d{0,9}\.\p{N}\d{0,9}`}

	assert.True(t, m.isProtectedBranches("master"))
	assert.True(t, m.isProtectedBranches("1.1.1"))
	assert.True(t, m.isProtectedBranches("2.0.0"))
	assert.False(t, m.isProtectedBranches("test"))
	assert.False(t, m.isProtectedBranches("hello1"))
}
