package gazzettabot

import (
	"log"
	"os"
	"os/exec"
)

type CliXdccBridge struct{}

func (this CliXdccBridge) Search(query string) []IrcFile {
	log.Printf("Performing search for query [%s] ...", query)

	var xdccBinary, xdccBinaryDefined = os.LookupEnv("XDCC_BINARY")
	if(!xdccBinaryDefined) {
		log.Fatal("XDCC_BINARY environment variable not defined!")
	}

	var cmd = exec.Command(xdccBinary, "search", query)
	var out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var outputTable = string(out)
	log.Println(outputTable)
	return []IrcFile{}
}
