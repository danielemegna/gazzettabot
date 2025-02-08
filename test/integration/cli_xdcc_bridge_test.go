package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchWithoutResults(t *testing.T) {
	var bridge = CliXdccBridge{}
	var files = bridge.Search("impossibile search string 19091990")
	assert.Empty(t, files)
	assert.Equal(t, []IrcFile{}, files)
}

func TestSearchWithSomeResults(t *testing.T) {
	var bridge = CliXdccBridge{}
	var files = bridge.Search("gazzetta dello sport completa 8 febbraio")
	assert.NotEmpty(t, files)
	assert.Contains(t, files[0].Name, "Gazzetta.dello.Sport")
	assert.Greater(t, files[0].SizeInMegaByte, 0)
	assert.Contains(t, files[0].Url, "irc://")
}
