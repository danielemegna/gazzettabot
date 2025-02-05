package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
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
	log.Println(fmt.Sprintf("%s", out))
}
