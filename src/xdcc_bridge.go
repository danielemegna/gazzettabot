package gazzettabot

import (
	"github.com/samber/lo"
	"log"
	"strconv"
	"strings"
)

type XdccBridge interface {
	Search(query string) []IrcFile
	Download(ircFileUrl string)
}

type IrcFile struct {
	Name           string
	SizeInMegaByte int
	Url            string
}

func IrcFilesToString(files []IrcFile) string {
	var filesToStrings = lo.Map(files, func(file IrcFile, _ int) string {
		return "  " + file.Name + " [" + strconv.Itoa(file.SizeInMegaByte) + "MB - " + file.Url + "]\n"
	})
	return "[\n" + strings.Join(filesToStrings, "") + "]"
}

func SmallestFrom(files []IrcFile) IrcFile {
	log.Println("Taking smallest from: " + IrcFilesToString(files))
	return lo.MinBy(files, func(a IrcFile, b IrcFile) bool {
		return a.SizeInMegaByte < b.SizeInMegaByte
	})
}
