package gazzettabot

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ParseTable(tableString string) []IrcFile {
	var rows = strings.Split(tableString, "\n")
	if len(rows) < 5 {
		return []IrcFile{}
	}

	var rowsWithoutHeaders = rows[3:]
	var foundFiles = rowsWithoutHeaders[:len(rowsWithoutHeaders)-2]
	return lo.Map(foundFiles, func(row string, _ int) IrcFile {
		var r, _ = regexp.Compile(`\|\s*(\S+)\s*\|\s*(\d+)MB\s*\|\s+(\S+)\s+\|`)
		var matches = r.FindStringSubmatch(row)
		var sizeInMegaByte, _ = strconv.Atoi(matches[2])
		return IrcFile{
			Name:           matches[1],
			SizeInMegaByte: sizeInMegaByte,
			Url:            matches[3],
		}
	})
}
