package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"github.com/stretchr/testify/assert"
	"testing"
)

var prioritizer = IrcFilePrioritizer{}

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

func TestFiftyShadesOfCompletaNonLocale(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.Ed..COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 15},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.definitiva.pdf", SizeInMegaByte: 30},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.definitiva.pdf", SizeInMegaByte: 30},
		{Name: "La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.Ed..COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 15},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
	}
	assert.Equal(t, expected, actual)
}

func TestPrioritizeLombardiaAsEdLocaleDespiteSize(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 100},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 100},
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf"},
	}
	assert.Equal(t, expected, actual)
}

func TestPrioritizeCompletaLombardiaOnNoLabelLombardia(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.pdf"},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf"},
	}
	assert.Equal(t, expected, actual)
}

func TestPrioritizeNoLabelLombardiaOnProvvisoriaLombardia(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.pdf"},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.versione.provvisoria.pdf"},
	}
	assert.Equal(t, expected, actual)
}

func TestProvvisoriePrioritization(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.21.Febbraio.2025.versione.provvisoria.pdf"},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.21.Febbraio.2025.versione.provvisoria.pdf"},
	}
	assert.Equal(t, expected, actual)
}

func TestEdCompletaShouldBePrioritized(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.Ed..COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf"},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed..COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf"},
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.febbraio.2025.pdf"},
	}
	assert.Equal(t, expected, actual)
}

func TestSortNonLombardiaEdLocaliBySize(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 40},
		{Name: "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.Ed.Sicilia.e.Calabria.21.Febbraio.2025.pdf", SizeInMegaByte: 30},
		{Name: "La.Gazzetta.dello.Sport.Ed.Verona.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Ed.Verona.21.Febbraio.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.Ed.Sicilia.e.Calabria.21.Febbraio.2025.pdf", SizeInMegaByte: 30},
		{Name: "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.Febbraio.2025.pdf", SizeInMegaByte: 40},
	}
	assert.Equal(t, expected, actual)
}

func TestCompletaWithEdizioniLocaliRealCaseWithoutEdKeyword(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Verona.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 40},
		{Name: "La.Gazzetta.dello.Sport.Lombardia.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 20},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 40},
		{Name: "La.Gazzetta.dello.Sport.Lombardia.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 20},
		{Name: "La.Gazzetta.dello.Sport.Verona.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 10},
	}
	assert.Equal(t, expected, actual)
}

func TestEdizioniLocaliRealCaseCompletaAndWithoutEdKeyword(t *testing.T) {
	var files = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Roma.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 15},
		{Name: "La.Gazzetta.dello.Sport.Bologna.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 11},
		{Name: "La.Gazzetta.dello.Sport.Cagliari.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 13},
		{Name: "La.Gazzetta.dello.Sport.Verona.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 12},
		{Name: "La.Gazzetta.dello.Sport.Lombardia.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 40},
		{Name: "La.Gazzetta.dello.Sport.Puglia.e.Basilicata.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 14},
		{Name: "La.Gazzetta.dello.Sport.Sicilia.e.Calabria.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 10},
	}

	var actual = prioritizer.SortGazzettaFiles(files)

	var expected = []IrcFile{
		{Name: "La.Gazzetta.dello.Sport.Lombardia.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 40},
		{Name: "La.Gazzetta.dello.Sport.Sicilia.e.Calabria.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 10},
		{Name: "La.Gazzetta.dello.Sport.Bologna.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 11},
		{Name: "La.Gazzetta.dello.Sport.Verona.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 12},
		{Name: "La.Gazzetta.dello.Sport.Cagliari.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 13},
		{Name: "La.Gazzetta.dello.Sport.Puglia.e.Basilicata.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 14},
		{Name: "La.Gazzetta.dello.Sport.Roma.COMPLETA.2.Marzo.2025.pdf", SizeInMegaByte: 15},
	}
	assert.Equal(t, expected, actual)
}

/*
	Cases we faced:
	+La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.definitiva.pdf
	+La.Gazzetta.dello.Sport.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf
	+La.Gazzetta.dello.Sport.Ed..COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.21.febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.COMPLETA.21.Febbraio.2025.versione.provvisoria.pdf
	+La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.21.febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Lombardia.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Lombardia.COMPLETA.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Sicilia.e.Calabria.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Ed.Verona.21.Febbraio.2025.pdf
	+La.Gazzetta.dello.Sport.Roma.COMPLETA.2.Marzo.2025.pdf
	+La.Gazzetta.dello.Sport.Bologna.COMPLETA.2.Marzo.2025.pdf
	+La.Gazzetta.dello.Sport.Verona.COMPLETA.2.Marzo.2025.pdf
	+La.Gazzetta.dello.Sport.Cagliari.COMPLETA.2.Marzo.2025.pdf
	+La.Gazzetta.dello.Sport.Lombardia.COMPLETA.2.Marzo.2025.pdf
	+La.Gazzetta.dello.Sport.Puglia.e.Basilicata.COMPLETA.2.Marzo.2025.pdf
	+La.Gazzetta.dello.Sport.Sicilia.e.Calabria.COMPLETA.2.Marzo.2025.pdf
*/
