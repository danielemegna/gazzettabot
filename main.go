package main

import (
	. "danielemegna/gazzettabot/src"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

var XDCC_BINARY_FILEPATH = getFromEnv("XDCC_BINARY")
var DOWNLOAD_FOLDER_PATH = getFromEnv("DOWNLOAD_FOLDER")

var xdccBridge XdccBridge = CliXdccBridge{
	XdccBinaryFilepath: XDCC_BINARY_FILEPATH,
	DownloadFolderPath: DOWNLOAD_FOLDER_PATH,
}
var alreadyDownloadedFilesProvider = FileSystemAlreadyDownloadedFilesProvider{
	DownloadFolderPath: DOWNLOAD_FOLDER_PATH,
}
var ircFilePrioritizer = IrcFilePrioritizer{}

func main() {
	log.Println("==== Starting Gazzetta Bot")

	var todayDay = time.Now().Day()
	var searchQuery = "Gazzetta dello Sport " + strconv.Itoa(todayDay) + " Febbraio -" + generateTimestampID()
	var foundFiles = xdccBridge.Search(searchQuery)

	var noAlreadyDownloaded = filterAlreadyDownloadedFiles(foundFiles)
	if len(noAlreadyDownloaded) > 0 {
		var prioritizedFiles = ircFilePrioritizer.SortGazzettaFiles(noAlreadyDownloaded)
		var fileToDownload = prioritizedFiles[0]
		log.Println("File selected for download: " + fileToDownload.Name)
		xdccBridge.Download(fileToDownload.Url)
	} else {
		log.Println("Cannot find new files to download!")
	}

	log.Println("==== Closing Gazzetta Bot")
}

func filterAlreadyDownloadedFiles(files []IrcFile) []IrcFile {
	var alreadyDownloadedFilenames = alreadyDownloadedFilesProvider.List()
	return lo.Filter(files, func(file IrcFile, _ int) bool {
		return !slices.Contains(alreadyDownloadedFilenames, file.Name)
	})
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
