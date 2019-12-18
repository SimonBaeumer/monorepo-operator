#!/usr/bin/env bash
set -e

docker-compose down || true
docker-compose up --build -d
docker exec -it integration_git-client_1 "./../suite/test.sh"