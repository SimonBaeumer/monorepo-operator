# monorepo-operator

A tool for managing monolithic repositories with subtree splits.

## Quick start

```bash
$ monorepo-operator init git@github.com:UserName path/to/repos
```

### Sync

`sync` the current branch to a target branch on the remote subtree repositories.

```bash
$ monorepo-operator sync [branch-name]
```

### Exec

`exec` executes shell commands on all projects.

```bash
$ monorepo-operator exec "echo hello"
> Execute on project01
hello
> Execute on project02
hello
```

### Project

`project` lets you execute some tasks against single projects.

### exec

`exec` executes shell commands on a project.

```bash
$ monorepo-operator project exec project01 "echo pwd"
> Execute on project01
/tmp/monorepo/.git/.subtree-repos/project01
```

### split

`split` creates a subtree split of the project and returns the hash of it.

```bash
$ monorepo-operator project split project01
44a603d1720dee64e8c4f5b13f5b5f2e87d54402
```

### Configuration

```yaml
# Mapping of projects to path inside the mono-repo and the corresponding git-url
projects:
- name: project
  path: projects/project01
  git-url: git@github.com:UserName/project02

- name: project
  path: project/project02
  git-url: git@github.com:UserName/project02

# operating-directory stores the original repositories with the git configs
# the exec command executes all commands on all directories located under the operating dir
operating-directory: .git/.subtree-repos
```