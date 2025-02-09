#!/bin/sh

set -x
SCRIPT_DIR=$(dirname -- "$( readlink -f -- "$0"; )";)

git pull --rebase --autostash
docker run --rm -v $SCRIPT_DIR:/app -w /app golang:1.23.3-alpine go build -o bin/main main.go
