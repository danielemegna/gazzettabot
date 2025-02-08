package main

import (
	. "danielemegna/gazzettabot/src"
	"github.com/samber/lo"
	"log"
	"os/exec"
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

	var noProvvisorie = lo.Filter(files, func(file IrcFile, _ int) bool {
		return !strings.Contains(file.Name, "provvisoria")
	})
	if len(noProvvisorie) == 0 {
		return smallest(files)
	}

	var noEdizioniLocali = lo.Filter(noProvvisorie, func(file IrcFile, _ int) bool {
		return !strings.Contains(file.Name, "Ed")
	})
	if len(noEdizioniLocali) == 0 {
		return smallest(noProvvisorie)
	}

	return smallest(noEdizioniLocali)
}

func smallest(files []IrcFile) IrcFile {
	return lo.MinBy(files, func(a IrcFile, b IrcFile) bool {
		return a.SizeInMegaByte < b.SizeInMegaByte // or > ?
	})
}
