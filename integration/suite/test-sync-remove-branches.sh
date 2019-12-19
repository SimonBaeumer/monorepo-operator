#!/usr/bin/env bash
set -e

monorepo-operator exec "git checkout -b not-exist-branch"
monorepo-operator exec "git push origin not-exist-branch"

monorepo-operator exec "git checkout -b not-exist-branch-local"
monorepo-operator exec "git checkout master"

echo "### CHECKING CURRENT LOCAL BRANCHES"
monorepo-operator exec "git branch" | diff./../suite/test-remove-branches-results.txt -

monorepo-operator remove-branches

echo "### CHECKING REMOVED BRANCHES"
monorepo-operator exec "git branch" | diff./../suite/test-remove-branches-results2.txt -
