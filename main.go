package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	log.Println("==== Starting Gazzetta Bot")
	var cmd = exec.Command("./lib/xdcc", "search", "gazzetta", "sport", "4", "febbraio")
	var out, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("%s", out))
}
