#!/usr/bin/env bash
set -e

function create_commits {
    for dir in subtrees/*/
    do
        dir=${dir%*/} #remove trailing /
        echo "test $dir" >> "$dir/README.md"
        git add "$dir/README.md"
        git commit -m "commit $dir"
    done
}

# create a new individual commit for every subtree
create_commits

monorepo-operator sync master
monorepo-operator exec "git fetch origin && git pull origin master"
