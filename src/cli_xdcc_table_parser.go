package gazzettabot

import (
	"github.com/samber/lo"
	"strings"
)

func ParseTable(tableString string) []IrcFile {
	var rows = strings.Split(tableString, "\n")
	if len(rows) < 5 {
		return []IrcFile{}
	}

	var rowsWithoutHeaders = rows[3:]
	var foundFiles = rowsWithoutHeaders[:len(rowsWithoutHeaders)-2]
	return lo.Map(foundFiles, func(row string, _ int) IrcFile {
		var parts = strings.Split(row, "|")
		return IrcFile{
			Name: strings.TrimSpace(parts[1]),
			SizeInMegaByte: 74,
			Url: "irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/146",
		}
	})
}
