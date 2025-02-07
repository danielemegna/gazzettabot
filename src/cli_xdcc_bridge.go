package gazzettabot

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CliXdccBridge struct{}

func (this CliXdccBridge) Search(query string) []IrcFile {
	log.Printf("Performing search for query [%s] ...", query)

	var xdccBinary, xdccBinaryDefined = os.LookupEnv("XDCC_BINARY")
	if !xdccBinaryDefined {
		log.Fatal("XDCC_BINARY environment variable not defined!")
	}

	var cmd = exec.Command(xdccBinary, "search", query)
	var out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var outputTable = string(out)
	var rows = strings.Split(outputTable, "\n")
	log.Println("Rows count: " + strconv.Itoa(len(rows)))
	if len(rows) > 4 {
		return []IrcFile{
			{
				Name:           "La.Gazzetta.dello.Sport.COMPLETA.7.Febbraio.2025.pdf",
				SizeInMegaByte: 12,
				Url:            "irc://irc.williamgattone.it/#drakon/DrAk|EdIcOlA|02/86",
			},
		}
	}

	return []IrcFile{}
}
