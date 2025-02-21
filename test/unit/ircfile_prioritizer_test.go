package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"github.com/stretchr/testify/assert"
	"testing"
)

var stubAlreadyDownloadedFilesProvider = StubAlreadyDownloadedFilesProvider{}
var prioritizer = IrcFilePrioritizer{
	AlreadyDownloadedFilesProvider: &stubAlreadyDownloadedFilesProvider,
}

func TestNoFileToPrioritize(t *testing.T) {
	var actual = prioritizer.SortGazzettaFiles([]IrcFile{})
	assert.Equal(t, []IrcFile{}, actual)
}

func TestSortBySizeWithSameFilename(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 21},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 14},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 14},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 21},
	}
	assert.Equal(t, expected, actual)
}

func TestPrioritizeSmallestCompletaOnProvvisoria(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 16},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 16},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 10},
	}
	assert.Equal(t, expected, actual)
}

func TestNoFileWhenEverythingAlreadyDownloaded(t *testing.T) {
	stubAlreadyDownloadedFilesProvider.SetAlreadyDownloadedFiles([]string{
		"La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf",
		"La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf",
	})
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 16},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 30},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	assert.Equal(t, []IrcFile{}, actual)
	stubAlreadyDownloadedFilesProvider.SetAlreadyDownloadedFiles([]string{})
}

func TestProvvisoriaWhenCompletedAlreadyDownloadedAndNoOtherChoice(t *testing.T) {
	stubAlreadyDownloadedFilesProvider.SetAlreadyDownloadedFiles([]string{
		"La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf",
	})
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 16},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 30},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 20},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 30},
	}
	assert.Equal(t, expected, actual)
	stubAlreadyDownloadedFilesProvider.SetAlreadyDownloadedFiles([]string{})
}

func TestPrioritizeProvvisoriaOnEdLocaliAndKeepBothSortedBySize(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 14},
		{Name: "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 25},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 20},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf", SizeInMegaByte: 25},
		{Name: "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 14},
	}
	assert.Equal(t, expected, actual)
}

/* TEST CASES
 * filter already downloaded file (use the collaborator)
 * prioritize Ed.Lombardia when no complete
 * prioritize Ed.Lombardia complete on Ed.Lombardia nolabel
 * prioritize Ed.Lombardia nolable on Ed.Lombardia provvisoria
 * prioritize lombardia ed locale on only no complete files
 */

/*
	Cases we faced:
	+La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.definitiva.pdf
	+La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf
	+La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf
	+La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Sicilia.e.Calabria.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Verona.21.Febbraio.2025.pdf
*/

type StubAlreadyDownloadedFilesProvider struct{ alreadyDownloadedFiles []string }

func (this StubAlreadyDownloadedFilesProvider) List() []string { return this.alreadyDownloadedFiles }
func (this *StubAlreadyDownloadedFilesProvider) SetAlreadyDownloadedFiles(files []string) {
	this.alreadyDownloadedFiles = files
}
