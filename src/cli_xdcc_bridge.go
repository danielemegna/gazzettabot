package gazzettabot

import (
	"log"
	"os"
	"os/exec"
)

type CliXdccBridge struct{}

func (this CliXdccBridge) Search(query string) []IrcFile {
	log.Printf("Performing search for query [%s] ...\n", query)
	var outputString = execSearch(query)
	var files = ParseTable(outputString)
	log.Printf("Found %d files!\n", len(files))
	return files
}

func (this CliXdccBridge) Download(ircFileUrl string) {
	log.Println("Downloading file " + ircFileUrl + " ...")
	var xdccBinaryFilepath = xdccBinaryFilepathFromEnv()
	var cmd = exec.Command(xdccBinaryFilepath, "get", ircFileUrl, "-o", "./download")
	var output, err = cmd.Output()
	if err != nil {
		log.Fatal("Error during file download! - ", err, " - ", string(output))
	}

	log.Println("Download completed!")
}

func execSearch(query string) string {
	var xdccBinaryFilepath = xdccBinaryFilepathFromEnv()
	var command = exec.Command(xdccBinaryFilepath, "search", query)
	var out, err = command.Output()
	var commandOutput = string(out)
	if err != nil {
		log.Fatal(err, " - ", commandOutput)
	}
	return commandOutput
}

func xdccBinaryFilepathFromEnv() string {
	var xdccBinaryFilepath, xdccBinaryDefined = os.LookupEnv("XDCC_BINARY")
	if !xdccBinaryDefined {
		log.Fatal("XDCC_BINARY environment variable not defined!")
	}
	return xdccBinaryFilepath
}
