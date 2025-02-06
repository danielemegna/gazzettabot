package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchWithoutResults(t *testing.T) {
	var bridge = CliXdccBridge{}
	var files = bridge.Search("impossibile search string 19091990")
	assert.Empty(t, files)
	assert.Equal(t, []IrcFile{}, files)
}
