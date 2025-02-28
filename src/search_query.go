package gazzettabot

import (
	"fmt"
	"strconv"
	"time"
)

func GazzettaDelloSportSearchQueryFor(date time.Time) string {
	var todayDay = date.Day()
	return "Gazzetta dello Sport " + strconv.Itoa(todayDay) + " Febbraio" + " -" + generateTimestampID()
}

func generateTimestampID() string {
	var unixNano = fmt.Sprintf("%d", time.Now().UnixNano())
	return unixNano[len(unixNano)-6:]
}
