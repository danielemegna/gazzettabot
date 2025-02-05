package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Println("==== Starting Gazzetta Bot")
	var todayDay = time.Now().Day()
	var cmd = exec.Command(
		"./lib/xdcc", "search", "gazzetta", "sport", strconv.Itoa(todayDay), "febbraio",
	)
	var out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var outputTable = string(out)
	var firstRow = strings.Split(outputTable, "\n")[3]
	log.Println(firstRow)

	var url = "irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/399"
	cmd = exec.Command(
		"./lib/xdcc", "get", url,
	)

	log.Println("Downloading " + url + " ....")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Download completed!")
}
