package gazzettabot

import (
	"github.com/samber/lo"
	"slices"
	"strings"
)

type IrcFilePrioritizer struct {
	AlreadyDownloadedFilesProvider
}

func (this IrcFilePrioritizer) SortGazzettaFiles(files []IrcFile) []IrcFile {
	var alreadyDownloadedFilenames = this.AlreadyDownloadedFilesProvider.List()
	var noAlreadyDownloaded = lo.Filter(files, func(file IrcFile, _ int) bool {
		return !slices.Contains(alreadyDownloadedFilenames, file.Name)
	})

	var toPrioritize, toUnderrate = chunkByPredicate(noAlreadyDownloaded, func(file IrcFile) bool {
		return !strings.Contains(file.Name, "provvisoria")
	})

	return append(sortBySize(toPrioritize), sortBySize(toUnderrate)...)
}

func chunkByPredicate(files []IrcFile, prioritizationPredicate func(IrcFile) bool) ([]IrcFile, []IrcFile) {
	var toPrioritize = []IrcFile{}
	var toUnderrate = []IrcFile{}
	for _, file := range files {
		if prioritizationPredicate(file) {
			toPrioritize = append(toPrioritize, file)
			continue
		}
		toUnderrate = append(toUnderrate, file)
	}

	return toPrioritize, toUnderrate
}

func sortBySize(files []IrcFile) []IrcFile {
	var cloneOfFiles = slices.Clone(files)
	slices.SortFunc(cloneOfFiles, func(a IrcFile, b IrcFile) int {
		return a.SizeInMegaByte - b.SizeInMegaByte
	})
	return cloneOfFiles
}
