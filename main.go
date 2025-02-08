package main

import (
	. "danielemegna/gazzettabot/src"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	log.Println("==== Starting Gazzetta Bot")
	var todayDay = time.Now().Day()
	var searchQuery = "gazzetta dello sport completa " + strconv.Itoa(todayDay) + " febbraio"

	var xdccBridge = CliXdccBridge{}
	var foundFiles = xdccBridge.Search(searchQuery)

	// TODO better select file to download
	var fileToDownload = foundFiles[0]

	log.Println("Downloading " + fileToDownload.Name + " ....")
	var cmd = exec.Command("./lib/xdcc", "get", fileToDownload.Url, "-o", "./download")
	var output, err = cmd.Output()
	if err != nil {
		log.Fatal(err, " - ", string(output))
	}

	log.Println("Download completed!")
}
