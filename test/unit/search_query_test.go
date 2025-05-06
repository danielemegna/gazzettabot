package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestGenerateWithKnownDate(t *testing.T) {
	var testCases = []struct {
		month time.Month
		day   int
		regex string
	}{
		{time.January, 22, `^Gazzetta dello Sport .22. Gennaio 2025 -\d{6}$`},
		{time.March, 14, `^Gazzetta dello Sport .14. Marzo 2025 -\d{6}$`},
		{time.June, 3, `^Gazzetta dello Sport .3. Giugno 2025 -\d{6}$`},
		{time.September, 19, `^Gazzetta dello Sport .19. Settembre 2025 -\d{6}$`},
		{time.December, 25, `^Gazzetta dello Sport .25. Dicembre 2025 -\d{6}$`},
	}

	for index, testCase := range testCases {
		t.Run("Test case #"+strconv.Itoa(index+1), func(t *testing.T) {
			var date = time.Date(2025, testCase.month, testCase.day, 6, 0, 0, 0, time.Local)
			var searchQuery = GazzettaDelloSportSearchQueryFor(date)
			assert.Regexp(t, testCase.regex, searchQuery)
		})
	}
}

func TestGenerateWithNowDate(t *testing.T) {
	var searchQuery = GazzettaDelloSportSearchQueryFor(time.Now())
	assert.Greater(t, len(searchQuery), 30)
	assert.Regexp(t, `^Gazzetta dello Sport .\d{1,2}. \S+ \d{4} -\d{6}$`, searchQuery)
}
