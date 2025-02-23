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

	if len(noAlreadyDownloaded) == 0 {
		return noAlreadyDownloaded
	}

	var predicates = []func(file IrcFile) bool{
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "definitiva") &&
				strings.Contains(name, "completa") &&
				!strings.Contains(name, "ed")
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "completa") &&
				!strings.Contains(name, "provvisoria") &&
				!strings.Contains(name, "ed")
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "ed.completa") || strings.Contains(name, "ed..completa")
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "completa") &&
				strings.Contains(name, "provvisoria") &&
				!strings.Contains(name, "ed")
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "provvisoria") && !strings.Contains(name, "ed")
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "lombardia") && strings.Contains(name, "completa")
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "lombardia") && !strings.Contains(name, "provvisoria")
		},
		func(file IrcFile) bool { return strings.Contains(strings.ToLower(file.Name), "lombardia") },
		func(file IrcFile) bool { return !strings.Contains(strings.ToLower(file.Name), "provvisoria") },
	}

	var prioritized = []IrcFile{}
	var rest = noAlreadyDownloaded
	for _, predicate := range predicates {
		var toPrioritize, toUnderrate = chunkByPredicate(rest, predicate)
		prioritized = append(prioritized, sortBySize(toPrioritize)...)
		rest = toUnderrate
	}

	prioritized = append(prioritized, rest...)
	return prioritized
}

func chunkByPredicate(files []IrcFile, prioritizationPredicate func(IrcFile) bool) ([]IrcFile, []IrcFile) {
	var matching = []IrcFile{}
	var nonMatching = []IrcFile{}
	for _, file := range files {
		if prioritizationPredicate(file) {
			matching = append(matching, file)
			continue
		}
		nonMatching = append(nonMatching, file)
	}

	return matching, nonMatching
}

func sortBySize(files []IrcFile) []IrcFile {
	var cloneOfFiles = slices.Clone(files)
	slices.SortFunc(cloneOfFiles, func(a IrcFile, b IrcFile) int {
		return a.SizeInMegaByte - b.SizeInMegaByte
	})
	return cloneOfFiles
}
