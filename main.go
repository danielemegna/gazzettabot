package main

import (
	. "danielemegna/gazzettabot/src"
	"github.com/samber/lo"
	"log"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Println("==== Starting Gazzetta Bot")
	var todayDay = time.Now().Day()
	var searchQuery = "gazzetta dello sport completa " + strconv.Itoa(todayDay) + " febbraio"

	var xdccBridge = CliXdccBridge{}
	var foundFiles = xdccBridge.Search(searchQuery)

	var fileToDownload = selectFileToDownload(foundFiles)
	log.Println("Downloading " + fileToDownload.Name + " ....")
	var cmd = exec.Command("./lib/xdcc", "get", fileToDownload.Url, "-o", "./download")
	var output, err = cmd.Output()
	if err != nil {
		log.Fatal(err, " - ", string(output))
	}

	log.Println("Download completed!")
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
	var entries, err = os.ReadDir("./download")
	if err != nil {
		log.Fatal("Error reading download folder! - ", err)
	}
	return lo.Map(entries, func(e os.DirEntry, _ int) string { return e.Name() })
}
