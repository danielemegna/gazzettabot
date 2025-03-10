package gazzettabot

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"
)

type CliXdccBridge struct {
	XdccBinaryFilepath string
	DownloadFolderPath string
}

func (this CliXdccBridge) Search(query string) []IrcFile {
	log.Printf("Performing search for query [%s] ...\n", query)
	var outputString = this.execSearch(query)
	var files = ParseTable(outputString)
	log.Printf("Found %d files!\n", len(files))
	return files
}

func (this CliXdccBridge) Download(ircFileUrl string) bool {
	log.Println("Downloading file " + ircFileUrl + " ...")

	var maxDurationInSeconds = 120
	var timeout = time.Duration(maxDurationInSeconds) * time.Second
	var contextWithTimeout, cancelFn = context.WithTimeout(context.Background(), timeout)
	defer cancelFn()

	var command = exec.CommandContext(
		contextWithTimeout,
		this.XdccBinaryFilepath, "get", ircFileUrl,
		"-o", this.DownloadFolderPath,
	)
	var out, err = command.Output()
	var commandOutput = string(out)
	if err != nil || isErrorOutput(commandOutput) {
		log.Println("Error during file download! - ", err)
		log.Println("Command output: ", commandOutput)
		return false
	}

	log.Println("Download completed!")
	return true
}

func (this CliXdccBridge) execSearch(query string) string {
	var command = exec.Command(this.XdccBinaryFilepath, "search", query)
	var out, err = command.Output()
	var commandOutput = string(out)
	if err != nil {
		log.Fatal(err, " - ", commandOutput)
	}
	return commandOutput
}

func isErrorOutput(output string) bool {
	return strings.Contains(output, "no valid irc url") ||
		strings.Contains(output, "invalid syntax")
}
