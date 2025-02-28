package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateWithNowDate(t *testing.T) {
	var searchQuery = GazzettaDelloSportSearchQueryFor(time.Now())
	assert.Greater(t, len(searchQuery), 15)
	assert.Regexp(t, `^Gazzetta dello Sport \d{2} \S+ -\d{6}$`, searchQuery)
}
