package gazzettabot

import (
	"github.com/samber/lo"
	"regexp"
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
		var r, _ = regexp.Compile(`\|\s*(\S+)\s*\|\s*(\S+)\s*\|\s+(\S+)\s+\|`)
		var matches = r.FindStringSubmatch(row)
		return IrcFile{
			Name:           matches[1],
			SizeInMegaByte: 74,
			Url:            matches[3],
		}
	})
}
