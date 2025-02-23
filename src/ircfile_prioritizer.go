package gazzettabot

import (
	"log"
	"slices"
	"strings"
)

type IrcFilePrioritizer struct {
	AlreadyDownloadedFilesProvider
}

func (this IrcFilePrioritizer) SortGazzettaFiles(files []IrcFile) []IrcFile {
	log.Printf("Sorting %d found files ...", len(files))
	var predicatesByImportance = predicatesByImportance()
	return sortByPredicates(files, predicatesByImportance)
}

func predicatesByImportance() []func(file IrcFile) bool {
	var isDefinitiva = func(fileName string) bool { return strings.Contains(fileName, "definitiva") }
	var isCompleta = func(fileName string) bool { return strings.Contains(fileName, "completa") }
	var isProvvisoria = func(fileName string) bool { return strings.Contains(fileName, "provvisoria") }
	var isEdizioneLocaleLombardia = func(fileName string) bool { return strings.Contains(fileName, "lombardia") }
	var isEdizioneLocale = func(fileName string) bool {
		return strings.Contains(fileName, "ed") &&
			!strings.Contains(fileName, "ed.completa") &&
			!strings.Contains(fileName, "ed..completa")
	}

	return []func(file IrcFile) bool{
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return isDefinitiva(name) && isCompleta(name) && !isEdizioneLocale(name)
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return strings.Contains(name, "completa") && !isProvvisoria(name) && !isEdizioneLocale(name)
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return isCompleta(name) && isProvvisoria(name) && !isEdizioneLocale(name)
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return isProvvisoria(name) && !isEdizioneLocale(name)
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return isEdizioneLocaleLombardia(name) && isCompleta(name)
		},
		func(file IrcFile) bool {
			var name = strings.ToLower(file.Name)
			return isEdizioneLocaleLombardia(name) && !isProvvisoria(name)
		},
		func(file IrcFile) bool { return isEdizioneLocaleLombardia(strings.ToLower(file.Name)) },
		func(file IrcFile) bool { return !isProvvisoria(strings.ToLower(file.Name)) },
	}
}

func sortByPredicates(rest []IrcFile, predicates []func(file IrcFile) bool) []IrcFile {
	var prioritized = []IrcFile{}
	for _, predicate := range predicates {
		var toPrioritize, toUnderrate = chunkByPredicate(rest, predicate)
		prioritized = append(prioritized, sortBySize(toPrioritize)...)
		rest = toUnderrate
	}

	return append(prioritized, rest...)
}

func chunkByPredicate(files []IrcFile, predicate func(IrcFile) bool) ([]IrcFile, []IrcFile) {
	var matching = []IrcFile{}
	var nonMatching = []IrcFile{}
	for _, file := range files {
		if predicate(file) {
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
