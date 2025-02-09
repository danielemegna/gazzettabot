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
	execDownload(ircFileUrl)
	log.Println("Download completed!")
}

func execDownload(ircFileUrl string) {
	var xdccBinaryFilepath = xdccBinaryFilepathFromEnv()
	var downloadFolderPath = downloadFolderPathFromEnv()
	var cmd = exec.Command(xdccBinaryFilepath, "get", ircFileUrl, "-o", downloadFolderPath)
	var output, err = cmd.Output()
	if err != nil {
		log.Fatal("Error during file download! - ", err, " - ", string(output))
	}
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

func xdccBinaryFilepathFromEnv() string { return getFromEnv("XDCC_BINARY") }
func downloadFolderPathFromEnv() string { return getFromEnv("DOWNLOAD_FOLDER") }
func getFromEnv(varName string) string {
	var value, defined = os.LookupEnv(varName)
	if !defined {
		log.Fatal(varName + " environment variable not defined!")
	}
	return value
}
