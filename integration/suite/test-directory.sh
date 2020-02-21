#!/usr/bin/env bash
set -e

cat > .monorepo-operator-dir.yml <<EOL
projects:
- name: "{{.DirName}}"
  path: subtrees
  git-url: root@git-server:/srv/git/{{.DirName}}.git
  is-dir: true
EOL

echo "### CHECKING LIST DIRECTORY OUTPUT"
monorepo-operator --config .monorepo-operator-dir.yml list | diff ./../suite/test-directory-results.txt -

echo "sync with directory scan"
monorepo-operator --config .monorepo-operator-dir.yml sync -f master

echo "fetch results"
monorepo-operator --config .monorepo-operator-dir.yml clone --reset

echo "### CHECKING SYNC DIRECTORY RESULT"
monorepo-operator --config .monorepo-operator-dir.yml exec ls | diff ./../suite/test-directory-sync-results.txt -
