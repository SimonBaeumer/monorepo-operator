#!/bin/sh

REPO_NAME=$1

mkdir "${REPO_NAME}.git"
cd ${REPO_NAME}.git

git init --shared=true
echo "${REPO_NAME}" >> README.md

git add .
git commit -m "initial commit"