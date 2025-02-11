package main

import (
	. "danielemegna/gazzettabot/src"
	"github.com/samber/lo"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var XDCC_BINARY_FILEPATH = getFromEnv("XDCC_BINARY")
var DOWNLOAD_FOLDER_PATH = getFromEnv("DOWNLOAD_FOLDER")

var xdccBridge XdccBridge = CliXdccBridge{
	XdccBinaryFilepath: XDCC_BINARY_FILEPATH,
	DownloadFolderPath: DOWNLOAD_FOLDER_PATH,
}
var alreadyDownloadedFilesProvider AlreadyDownloadedFilesProvider = FileSystemAlreadyDownloadedFilesProvider{
	DownloadFolderPath: DOWNLOAD_FOLDER_PATH,
}

func main() {
	log.Println("==== Starting Gazzetta Bot")

	var todayDay = time.Now().Day()
	var searchQuery = "GAZZETTA DELLO SPORT " + strconv.Itoa(todayDay) + " FEBBRAIO"
	var foundFiles = xdccBridge.Search(searchQuery)

	var alreadyDownloadedFilenames = alreadyDownloadedFilesProvider.List()
	var fileToDownload = selectFileToDownload(foundFiles, alreadyDownloadedFilenames)
	log.Println("File selected for download: " + fileToDownload.Name)

	xdccBridge.Download(fileToDownload.Url)
	log.Println("==== Closing Gazzetta Bot")
}

func selectFileToDownload(files []IrcFile, alreadyDownloadedFilenames []string) IrcFile {
	if len(files) == 0 {
		log.Fatal("Gazzetta not Found !")
	}
	if len(files) == 1 {
		return files[0]
	}

	var noAlreadyDownloaded = lo.Filter(files, func(file IrcFile, _ int) bool {
		return !slices.Contains(alreadyDownloadedFilenames, file.Name)
	})
	if len(noAlreadyDownloaded) == 0 {
		log.Fatal("Cannot find new file to download!")
	}

	var noProvvisorie = lo.Filter(noAlreadyDownloaded, func(file IrcFile, _ int) bool {
		return !strings.Contains(file.Name, "provvisoria")
	})
	if len(noProvvisorie) == 0 {
		return SmallestFrom(noAlreadyDownloaded)
	}

	var onlyComplete = lo.Filter(noProvvisorie, func(file IrcFile, _ int) bool {
		return strings.Contains(strings.ToLower(file.Name), "completa")
	})
	if len(onlyComplete) == 0 {
		return SmallestFrom(noProvvisorie)
	}

	var noEdizioniLocali = lo.Filter(onlyComplete, func(file IrcFile, _ int) bool {
		return !strings.Contains(strings.ToLower(file.Name), "ed.")
	})

	if len(noEdizioniLocali) == 0 {
		var edLombardia = lo.Filter(onlyComplete, func(file IrcFile, _ int) bool {
			return strings.Contains(strings.ToLower(file.Name), "lombardia")
		})
		if len(edLombardia) > 0 {
			return SmallestFrom(edLombardia)
		} else {
			return SmallestFrom(onlyComplete)
		}
	}

	return SmallestFrom(noEdizioniLocali)
}

func getFromEnv(varName string) string {
	var value, defined = os.LookupEnv(varName)
	if !defined {
		log.Fatal(varName + " environment variable not defined!")
	}
	return value
}
