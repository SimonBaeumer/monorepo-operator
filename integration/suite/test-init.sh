#!/usr/bin/env bash
set -e

monorepo-operator init root@git-server:/srv/git/ subtrees/ --clone

git add .monorepo-operator.yml
git commit -m "Init monorepo-operator"
git push origin master

echo "## CHECKING GENERATED MONOREPO-OPERATOR.YML"
cat .monorepo-operator.yml | diff ./../suite/test-init-results.txt -
