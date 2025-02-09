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

var xdccBridge XdccBridge = CliXdccBridge{}

func main() {
	log.Println("==== Starting Gazzetta Bot")

	var todayDay = time.Now().Day()
	var searchQuery = "gazzetta dello sport completa " + strconv.Itoa(todayDay) + " febbraio"
	var foundFiles = xdccBridge.Search(searchQuery)

	var fileToDownload = selectFileToDownload(foundFiles)
	log.Println("File selected for download: " + fileToDownload.Name)

	xdccBridge.Download(fileToDownload.Url)
	log.Println("==== Closing Gazzetta Bot")
}

func selectFileToDownload(files []IrcFile) IrcFile {
	if len(files) == 0 {
		log.Fatal("Gazzetta not Found !")
	}
	if len(files) == 1 {
		return files[0]
	}

	var alreadyDownloaded = getAlreadyDownloadedFileNames()
	var noAlreadyDownloaded = lo.Filter(files, func(file IrcFile, _ int) bool {
		return !slices.Contains(alreadyDownloaded, file.Name)
	})
	if len(noAlreadyDownloaded) == 0 {
		log.Fatal("Cannot find new file to download!")
	}

	var noProvvisorie = lo.Filter(noAlreadyDownloaded, func(file IrcFile, _ int) bool {
		return !strings.Contains(file.Name, "provvisoria")
	})
	if len(noProvvisorie) == 0 {
		return smallest(noAlreadyDownloaded)
	}

	var noEdizioniLocali = lo.Filter(noProvvisorie, func(file IrcFile, _ int) bool {
		return !strings.Contains(file.Name, "Ed")
	})

	if len(noEdizioniLocali) == 0 {
		var edLombardia = lo.Filter(noProvvisorie, func(file IrcFile, _ int) bool {
			return strings.Contains(file.Name, "Ed.Lombardia")
		})
		if len(edLombardia) > 0 {
			return smallest(edLombardia)
		} else {
			return smallest(noProvvisorie)
		}
	}

	return smallest(noEdizioniLocali)
}

func smallest(files []IrcFile) IrcFile {
	log.Println("Taking smallest from filtered: " + IrcFilesToString(files))
	return lo.MinBy(files, func(a IrcFile, b IrcFile) bool {
		return a.SizeInMegaByte < b.SizeInMegaByte // or > ?
	})
}

func getAlreadyDownloadedFileNames() []string {
	var entries, err = os.ReadDir(downloadFolderPathFromEnv())
	if err != nil {
		log.Fatal("Error reading download folder! - ", err)
	}
	return lo.Map(entries, func(e os.DirEntry, _ int) string { return e.Name() })
}

// TODO remove duplication of this snippet also in CliXdccBridge
func downloadFolderPathFromEnv() string { return getFromEnv("DOWNLOAD_FOLDER") }
func getFromEnv(varName string) string {
	var value, defined = os.LookupEnv(varName)
	if !defined {
		log.Fatal(varName + " environment variable not defined!")
	}
	return value
}