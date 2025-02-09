package gazzettabot

import (
	"log"
	"os/exec"
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

func (this CliXdccBridge) Download(ircFileUrl string) {
	log.Println("Downloading file " + ircFileUrl + " ...")
	this.execDownload(ircFileUrl)
	log.Println("Download completed!")
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

func (this CliXdccBridge) execDownload(ircFileUrl string) {
	var cmd = exec.Command(
		this.XdccBinaryFilepath, "get", ircFileUrl,
		"-o", this.DownloadFolderPath,
	)
	var output, err = cmd.Output()
	if err != nil {
		log.Fatal("Error during file download! - ", err, " - ", string(output))
	}
}
