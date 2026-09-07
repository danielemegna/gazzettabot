package gazzettabot

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type XdccBridge interface {
	Search(query string) []IrcFile
	DownloadOneOf(ircFiles []IrcFile)
}

type IrcFile struct {
	Name           string
	SizeInMegaByte float64
	Url            string
}

func IrcFilesToString(files []IrcFile) string {
	var filesToStrings = lo.Map(files, func(file IrcFile, _ int) string {
		return "  " + file.Name + " [" + strconv.FormatFloat(file.SizeInMegaByte, 'f', 2, 32) + "MB - " + file.Url + "]\n"
	})
	return "[\n" + strings.Join(filesToStrings, "") + "]"
}
