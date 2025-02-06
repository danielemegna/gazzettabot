package gazzettabot

import (
	"log"
	"os"
	"os/exec"
)

type CliXdccBridge struct{}

func (this CliXdccBridge) Search(query string) []IrcFile {
	log.Printf("Performing search for query [%s] ...", query)

	wd, _ := os.Getwd()
	log.Println(wd)

	var cmd = exec.Command("??./lib/xdcc", query)
	var out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var outputTable = string(out)
	log.Println(outputTable)
	return []IrcFile{}
}
