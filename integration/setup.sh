#!/bin/bash

mkdir monorepo
cd monorepo

git init
touch test
git add -A
git commit -m "Initial commit"

git remote add origin root@git-server:/srv/git/monorepo.git
git push origin master