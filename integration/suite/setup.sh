#!/bin/bash

git init
git add -A
git commit -m "Initial commit"

git remote add origin root@git-server:/srv/git/monorepo.git
git push origin master