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
- [x] use `completa` just as another prioritization filter (remove it from query)
- [x] optimize update sh script
- [x] prioritize lombardia ed locale on only no complete files
- [ ] month to italian string
- [ ] add year to search string
- [ ] try next on download failure: selectFileToDownload should return a prioritized list of files ?
- [ ] handle GigaByte as size in cli_xdcc_table_parser
- [ ] take another random file on many files with same name (already downloaded) ?
- [x] [refactor] decouple AlreadyDownloadedFilesProvider from IrcFilePrioritizer
- [x] [refactor] extract download operation in XdccBridge
- [x] [refactor] improve readability on selectFileToDownload: try filter & download on too much filter
- [x] [refactor] add CliXdccBridge integration tests for Download function
- [ ] [refactor] add FileSystemAlreadyDownloadedFilesProvider integration tests for List function
- [ ] [refactor] replace CliXdccBridge#search with an http implementation (direct fetch from xdcc.eu/search.php?searchkey=query)
- [x] [refactor] add unit tests for SmallestFrom([]IrcFile) function
- [x] [refactor] remove duplication on env variables reading: move in main and inject?
- [x] [refactor] download folder path as env variable
