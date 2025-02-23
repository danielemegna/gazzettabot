package main

import (
	. "danielemegna/gazzettabot/src"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var XDCC_BINARY_FILEPATH = getFromEnv("XDCC_BINARY")
var DOWNLOAD_FOLDER_PATH = getFromEnv("DOWNLOAD_FOLDER")

var xdccBridge XdccBridge = CliXdccBridge{
	XdccBinaryFilepath: XDCC_BINARY_FILEPATH,
	DownloadFolderPath: DOWNLOAD_FOLDER_PATH,
}
var ircFilePrioritizer = IrcFilePrioritizer{
	AlreadyDownloadedFilesProvider: FileSystemAlreadyDownloadedFilesProvider{
		DownloadFolderPath: DOWNLOAD_FOLDER_PATH,
	},
}

func main() {
	log.Println("==== Starting Gazzetta Bot")

	var todayDay = time.Now().Day()
	var searchQuery = "Gazzetta dello Sport " + strconv.Itoa(todayDay) + " Febbraio -" + generateTimestampID()
	var foundFiles = xdccBridge.Search(searchQuery)

	var prioritizedFiles = ircFilePrioritizer.SortGazzettaFiles(foundFiles)
	if len(prioritizedFiles) > 0 {
		var fileToDownload = prioritizedFiles[0]
		log.Println("File selected for download: " + fileToDownload.Name)
		xdccBridge.Download(fileToDownload.Url)
	}

	log.Println("==== Closing Gazzetta Bot")
}

func generateTimestampID() string {
	var unixNano = fmt.Sprintf("%d", time.Now().UnixNano())
	return unixNano[len(unixNano)-6:]
}

func getFromEnv(varName string) string {
	var value, defined = os.LookupEnv(varName)
	if !defined {
		log.Fatal(varName + " environment variable not defined!")
	}
	return value
}
