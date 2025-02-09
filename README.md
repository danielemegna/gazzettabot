## Go Gazzetta Bot

Build and run with:

```
$ go build -o bin/main main.go
$ ./bin/main
```

### Docker-only dev setup

Use temporary container with current directory volume:

```
$ docker run --rm -it -v $PWD:/app -w /app golang:1.23.3 bash
# go install github.com/mfridman/tparse@latest
```

and run tests with
```
# go test ./... -json | tparse
# go test ./.../unit -json | tparse --all
# go test ./.../integration -json | tparse --all
# go test ./... -run TestRunSingleTest -json | tparse --all
```

### VSCode Dev Container setup

VSCode can use devcontainers to be configured with the proper extensions without language specific utilities locally installed.

In order to proceed, install the [related extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) (`ms-vscode-remote.remote-containers`) and reopen vscode in Dev Container mode.

> the first run can take some minutes to install and setup properly container and vscode extensions

If you want to open an additional bash session in the vscode container:

```
$ docker exec -it -w /workspaces/$(basename $PWD) <container-name> bash
# go test .....
```

### On the fly docker build

Build the executable with a temporary alpine go container:

```
$ docker run --rm -v $PWD:/app -w /app golang:1.23.3-alpine go build -o bin/main main.go
```

### Backlog

- [x] print filtered files before select
- [x] prioritize Ed.Lombardia before others
- [ ] month to italian string
- [ ] add year to search string
- [ ] try next on download failure: selectFileToDownload should return a prioritized list of files ?
- [ ] take another random file on many files with same name (already downloaded)
- [x] [refactor] extract download operation in XdccBridge
- [ ] [refactor] improve readability on selectFileToDownload: try filter & download on too much filter
- [ ] [refactor] download folder path as env variable
