#!/usr/bin/env bash
set -e

monorepo-operator exec "git fetch origin && git pull origin master"

function create_commits {
    for dir in subtrees/*/
    do
        dir=${dir%*/} #remove trailing /
        echo "Release v1.0.0-test" >> "$dir/README.md"
        git add "$dir/README.md"
        git commit -m "Release $dir v1.0.0-test"
    done
}

create_commits

# Sync new commits
monorepo-operator sync master

# Update subtrees
monorepo-operator exec "git fetch origin && git pull origin master"

# Create tag
git tag -a v1.0.0-test -m v1.0.0-test

# Push tag
git push origin v1.0.0-test

# Sync tags
monorepo-operator sync v1.0.0-test --tags

# Fetch updates
monorepo-operator exec git fetch origin

echo "### CHECKING SYNCED TAGS"
monorepo-operator exec "git log --name-status --pretty=format:'%an, %ae - %s, NAMES: %D' HEAD^..HEAD" | diff ./../suite/test-sync-tag-results.txt -
