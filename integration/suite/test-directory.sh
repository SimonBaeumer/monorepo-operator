#!/usr/bin/env bash
set -e

cat > .monorepo-operator-dir.yml <<EOL
projects:
- name: "{{.DirName}}"
  path: subtrees
  git-url: root@git-server:/srv/git/{{.DirName}}.git
  is-dir: true
- name: manual-repo
  path: /tmp/manual
  git-url: root@git-server:/srv/git/manual.git
EOL

echo "### CHECKING LIST DIRECTORY OUTPUT"
monorepo-operator --config .monorepo-operator-dir.yml list | diff ./../suite/test-directory-results.txt -