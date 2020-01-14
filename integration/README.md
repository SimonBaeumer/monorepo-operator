# Integration tests

## Requirements

 - docker-compose 1.22
 - docker 19.03
 - bash

## Executing

Just start the integration script, it will setup a compose environment and executing the synchronisation.
During the execution it will do several assertions for validating the processes.

```$ ./integration.sh```

## Debbuging

### git-client

The git-client represents the client which will execute the sync and hold the different repos.

**Shell:**

`$ docker exec -it integration_git-client_1 /bin/bash`

**Executing tests**

```bash
$ ./../suite/test.sh # Execute complete suite on a clean environment
$ ./../suite/setup.sh # Execute only the setup of the test suite
$ ./../suite/test-init.sh # Execute the init test
$ ./../suite/test-sync.sh # Execute the sync test, currently depends on the init test
```

### git-server

The git-server is the remote git server where the monorepo should be synchronized into.
You can access it with `git@git-server:/srv/git/[repo-name].git`

Repos:

 - `git@git-server:/srv/git/monorepo.git`
 - `git@git-server:/srv/git/subtree1.git`
 - `git@git-server:/srv/git/subtree2.git`
 - `git@git-server:/srv/git/subtree3.git`
 
The repos are located in the `git-server` container in `/srv/git/`.

**Shell:**

`$ docker exec -it integration_git-client_1 /bin/bash`

## Test cases

 - [x] Init project
 - [ ] Split project
 - [x] Synchronize project
 - [x] Remove branches from subtrees which do not exist in monorepo
 - [x] Execute commands on every repository
 - [ ] Add a new project
 - [ ] Clone with `--reset`
 