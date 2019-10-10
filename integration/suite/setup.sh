#!/bin/sh

GIT_URL=$1

mkdir /monorepo.git/repo01
mkdir /monorepo.git/repo02

git init --shared=true
git add .
git commit -m "Initial commit"

scp -r /monorepo.git git@git-server:/git-server/repos

mkdir /testing
git clone ssh://git@git-server/git-server/repos/monorepo.git