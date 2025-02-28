package gazzettabot

import (
	"fmt"
	"strconv"
	"time"
)

func GazzettaDelloSportSearchQueryFor(date time.Time) string {
	var todayDayString = strconv.Itoa(date.Day())
	var todayMonthString = italianNameFor(date.Month())
	var dateString = todayDayString + " " + todayMonthString
	return "Gazzetta dello Sport " + dateString + " -" + generateTimestampID()
}

func italianNameFor(month time.Month) string {
	var lookup = []string{
		"Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno",
		"Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre",
	}
	return lookup[month-1]
}

func generateTimestampID() string {
	var unixNano = fmt.Sprintf("%d", time.Now().UnixNano())
	return unixNano[len(unixNano)-6:]
}
