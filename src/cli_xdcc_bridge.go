package gazzettabot

import (
	"log"
	"os"
	"os/exec"
)

type CliXdccBridge struct{}

func (this CliXdccBridge) Search(query string) []IrcFile {
	log.Printf("Performing search for query [%s] ...\n", query)
	var xdccBinaryFilepath = xdccBinaryFilepathFromEnv()
	var outputString = execSearch(xdccBinaryFilepath, query)
	var files = ParseTable(outputString)
	log.Printf("Found %d files!\n", len(files))
	return files
}

func execSearch(xdccBinaryFilepath string, query string) string {
	var command = exec.Command(xdccBinaryFilepath, "search", query)
	var commandOutput, err = command.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(commandOutput)
}

func xdccBinaryFilepathFromEnv() string {
	var xdccBinaryFilepath, xdccBinaryDefined = os.LookupEnv("XDCC_BINARY")
	if !xdccBinaryDefined {
		log.Fatal("XDCC_BINARY environment variable not defined!")
	}
	return xdccBinaryFilepath
}
