#!/usr/bin/env bash
set -e

rm -rf .git/.subtree-repos

monorepo-operator clone

echo "### CHECKING CLONE RESULTS"
ls -l .git/.subtree-repos/ | awk '{print $9}' | diff ./../suite/test-clone-results.txt -
