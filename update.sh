#!/bin/sh

set -x
SCRIPT_DIR=$(dirname -- "$( readlink -f -- "$0"; )";)

git pull --rebase --autostash

# ======== to run a slower unoptimized build from golang official image
#docker run --rm -v $SCRIPT_DIR:/app -w /app golang:1.23.3-alpine go build -o bin/main main.go

# ======== to create a gazzettabuild custom image to optimize following builds
#docker run --name gazzettabuild-container -v $SCRIPT_DIR:/app -w /app golang:1.23.3-alpine go build -o bin/main main.go
#docker commit gazzettabuild-container gazzettabuild && docker rm gazzettabuild-container && docker rmi golang:1.23.3-alpine

# ======== to build with the optimized gazzzettabuild custom image
docker run --rm -v $SCRIPT_DIR:/app -w /app gazzettabuild go build -o bin/main main.go
