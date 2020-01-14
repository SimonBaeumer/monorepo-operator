# monorepo-operator

A tool for managing monolithic repositories with subtree splits.


## Table of contents

* [Quick start](#quick-start)
  + [Requirements](#requirements)
  + [Usage](#usage)
    - [clone](#clone)
    - [sync](#sync)
    - [exec](#exec)
    - [add](#add)
    - [add](#remove-branches)
    - [project](#project)
      * [exec](#exec)
      * [split](#split)
  + [Configuration](#configuration)
* [Development](#development)
  + [Targets](#targets)
  + [ToDo](#todo)

## Quick start

This tool is for monolithic repository which need to synchronize directories into subtree repos.
A common use case is maintaining plugins or themes in the same repository.

```
plugins/
├── repo01
└── repo02
```

The `init` command create the necessary mappings for your directories. You can find this example in the 
[_examples](_examples) directory. 

The `--operating-dir` flag specifies the directory where the repositories related to the directories will be stored. 
These will be used for performing batch tasks on all subtrees with the `exec` command.

```bash
# Initialize the repos with the operating directory ".repos"
$ monorepo-operator init --operating-dir=.repos git@github.com:SimonBaeumer plugins/

# View written config file
$ cat .monorepo-operator.yml
projects:
- name: repo01
  path: plugins/repo01
  git-url: git@github.com:SimonBaeumer/repo01
- name: repo02
  path: plugins/repo02
  git-url: git@github.com:SimonBaeumer/repo02
operating-directory: .repos

# Clone the repository into the specified operating-directory .repos
$ monorepo-operator clone
> Cloning repo01
> Cloning repo02

$ ls -la .repos/
total 16
drwxrwxr-x 4 user user 4096 Okt 22 13:50 .
drwxrwxr-x 5 user user 4096 Okt 22 13:44 ..
drwxrwxr-x 3 user user 4096 Okt 22 13:50 repo01
drwxrwxr-x 3 user user 4096 Okt 22 13:50 repo02

# Execute a command on all subtree-repos directly
$ monorepo-operator exec "git remote -v"
> Execute on repo01
origin  git@github.com:SimonBaeumer/repo01.git (fetch)
origin  git@github.com:SimonBaeumer/repo01.git (push)
> Execute on repo02
origin  git@github.com:SimonBaeumer/repo02.git (fetch)
origin  git@github.com:SimonBaeumer/repo02.git (push)
```

They `sync` can only be performed from the root of your monorepo repo. This will not work with the example. 

`sync` creates subtrees for each project in the `.monorepo-operator.yml` with the 
`git subtree split` command. After that it pushed the changes to the desired repository, in this case `testing`.

```bash
# Create subtrees from currently checked out ref and push it to the configured repos on branch testing.
$ monorepo-sync tesing
> split project repo01 in branch repo01-testing
> add remote repo01
> push project repo01
Counting objects: 6, done.
[...]
To github.com:SimonBaeumer/repo01
 * [new branch]      repo01-testing -> testing
> remove remote repo01
> remove branch repo01-testing
Deleted branch repo01-testing (was a8f688f).

> split project repo02 in branch repo02-testing
> add remote repo02
> push project repo02
Total 0 (delta 0), reused 0 (delta 0)
[...]
To github.com:SimonBaeumer/repo02
 * [new branch]      repo02-testing -> testing
> remove remote repo02
> remove branch repo02-testing
Deleted branch repo02-testing (was 8fbf026).
``` 

### Requirements

 - git
 - `windows`, `osx` or `linux`
 
### Usage
 
#### clone

`clone` clones all projects into the specified `operating-directory`.

```bash
# If the operating-directory exists the command fails 
$ monorepo-operator clone
> Cloning repo01
2019/10/22 14:05:34 error while cloning: fatal: destination path '.git/.subtree-repos/repo01' already exists and is not an empty directory.

# Overwrite and re-clone all project with the --reset flag
$ monorepo-operator clone --reset
> Removing operating directory at .git/.subtree-repos
> Cloning repo01
> Cloning repo02
```

#### sync

`sync` the current branch to a target branch on the remote subtree repositories. 
This command only works in the root directory of your mono-repo.

If the `--force` flag is set the `sync` will perform a force push with `git push -f [...]`.

The `--remove-branches` flag removes branches in subtree repos which do not exist in the mono-repo.

The `--tags` flag syncs a tag instead of the given branch.  

```bash
# Sync branches
$ monorepo-operator sync [branch-name]

# Sync tags
$ monorepo-operator sync [tag-name] --tags
```

#### exec

`exec` executes shell commands on all projects.

```bash
$ monorepo-operator exec "echo hello"
> Execute on project01
hello
> Execute on project02
hello
```

#### add

`add` adds a new project mapping to the `.monorepo-operator.yml`.


```bash
# Add project to your mapping config
$ monorepo-operator add repo03 git@github.com:SimonBaeumer/repo03 repos/
> Write config file .monorepo-operator.yml

# Directly clone the repo of the project with --clone
$ monorepo-operator add --clone repo03 git@github.com:SimonBaeumer/repo03 repos/
> Write config file .monorepo-operator.yml
> Cloning repo03
[...]
```

#### remove-branches

`remove-branches` removes branches which do not exist locally or on the remote mono-repo in subtree repos.

`--no-local` disables removing local branches which do not exist in the remote mono repo.
`--no-remote` disables removing remote branches in subtree-repos which do not exist in the remote mono repo.

#### project

`project` lets you execute some commands or tasks on a single project.

##### exec

`project exec [name]` executes shell commands on a project.

```bash
$ monorepo-operator project exec repo01 "echo pwd"
> Execute on project01
/tmp/monorepo/.git/.subtree-repos/repo01
```

##### split

`project split [name]` creates a subtree split of the project and returns the hash of it.

```bash
$ monorepo-operator project split repo01
44a603d1720dee64e8c4f5b13f5b5f2e87d54402
```

### Configuration

```yaml
# Mapping of projects to path inside the mono-repo and the corresponding git-url
projects:
- name: project01
  path: projects/project01
  git-url: git@github.com:UserName/project02

- name: project02
  path: projects/project02
  git-url: git@github.com:UserName/project02

# operating-directory stores the original repositories with the git configs
# the exec command executes all commands on all directories located under the operating dir
operating-directory: .git/.subtree-repos
```

## Development

### Targets

```bash
# Init dev environment, i.e. git-hooks
$ make init

# Build project
$ make build

# Create releases
$ make release

# Execute unit tests
$ make tests
```

### ToDo

 - Lock and Unlock projects while executing commands
 - Post and Pre-Hooks
    - Split
    - Push
    - Exec
 - Add pipeline examples
