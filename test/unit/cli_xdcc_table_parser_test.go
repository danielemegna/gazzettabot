package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var EMPTY_TABLE_STRING = []string{
	"+------------------------+",
	"| File Name | Size | URL |",
	"+------------------------+",
}

var FULL_TABLE_STRING = []string{
	"+------------------------------------------------------------------------------------------------------------------------------------------------------+",
	"|                                File Name                                      | Size  |                            URL                               |",
	"+------------------------------------------------------------------------------------------------------------------------------------------------------+",
	"|       La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.3.Febbraio.2025.pdf         | 74MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/146    |",
	"|       La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.3.Febbraio.2025.pdf        | 74MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/147    |",
	"| La.Gazzetta.dello.Sport.Ed.Puglia.e.Basilicata.COMPLETA.3.Febbraio.2025.pdf   | 73MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/148    |",
	"|         La.Gazzetta.dello.Sport.Ed.Roma.COMPLETA.3.Febbraio.2025.pdf          | 73MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/149    |",
	"|  La.Gazzetta.dello.Sport.Ed.Sicilia.e.Calabria.COMPLETA.3.Febbraio.2025.pdf   | 74MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/150    |",
	"|       La.Gazzetta.dello.Sport.3.Febbraio.2025.versione.provvisoria.pdf        | 42MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/20     |",
	"|       La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.4.Febbraio.2025.pdf         | 62MB  |  irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/289    |",
	"+------------------------------------------------------------------------------------------------------------------------------------------------------+",
}

func TestParseEmptyTable(t *testing.T) {
	var files = ParseTable(asSingleString(EMPTY_TABLE_STRING))
	assert.Empty(t, files)
	assert.Equal(t, []IrcFile{}, files)
}

func TestParseFullTable(t *testing.T) {
	var files = ParseTable(asSingleString(FULL_TABLE_STRING))
	assert.Len(t, files, 7)
	assert.Equal(t, "La.Gazzetta.dello.Sport.Ed.Bologna.COMPLETA.3.Febbraio.2025.pdf", files[0].Name)
	assert.Equal(t, 74, files[0].SizeInMegaByte)
	assert.Equal(t, "irc://irc.arabaphenix.it/#arabafenice/ArA|Edicola|01/146", files[0].Url)
	assert.Equal(t, "La.Gazzetta.dello.Sport.Ed.Cagliari.COMPLETA.3.Febbraio.2025.pdf", files[1].Name)
}

func asSingleString(inputLines []string) string {
	return strings.Join(inputLines, "\n") + "\n"
}
