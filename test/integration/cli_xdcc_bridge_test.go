package gazzettabot_test

import (
	. "danielemegna/gazzettabot/src"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var bridge XdccBridge = CliXdccBridge{
	XdccBinaryFilepath: os.Getenv("XDCC_BINARY"),
	DownloadFolderPath: os.Getenv("DOWNLOAD_FOLDER"),
}

func TestSearchWithoutResults(t *testing.T) {
	var files = bridge.Search("impossibile search string 19091990")
	assert.Empty(t, files)
	assert.Equal(t, []IrcFile{}, files)
}

func TestSearchWithSomeResults(t *testing.T) {
	var files = bridge.Search("gazzetta dello sport completa 8 febbraio")
	assert.NotEmpty(t, files)
	assert.Contains(t, files[0].Name, "Gazzetta.dello.Sport")
	assert.Greater(t, files[0].SizeInMegaByte, 0)
	assert.Contains(t, files[0].Url, "irc://")
}

func TestDownloadFileWithSuccess(t *testing.T) {
	var ircFileUrl = "irc://irc.irc-files.org/#FaNtAsYlAnD/FaNtAsYlAnD|AnImE|01/3112"
	var expectedDownloadedFilename = "Star.Wars.Rebels.3x20.Ultimo.Atto.Prima.Parte.ITA.ENG.DLMux.XviD-Pir8.srt"

	var result = bridge.Download(ircFileUrl)

	assert.True(t, result, "Returned value should be true")
	var downloadFolder = os.Getenv("DOWNLOAD_FOLDER")
	var expectedDownloadedFilepath = downloadFolder + "/" + expectedDownloadedFilename
	var fileInfo, err = os.Stat(expectedDownloadedFilepath)
	assert.Nil(t, err, "Cannot find expected downloaded file")
	assert.Equal(t, int64(175), fileInfo.Size())
	err = os.Remove(downloadFolder + "/" + expectedDownloadedFilename)
	assert.Nil(t, err, "Cannot delete downloaded file")
}

func TestDownloadFileReturnsFalseOnTimeout(t *testing.T) {
	t.Skip("Very slow test, it use the timeout - todo move it in env var")
	var ircFileUrl = "irc://irc.openjoke.org/#William&Carola/WeC|EdIcOLa|01/3010"
	var result = bridge.Download(ircFileUrl)
	assert.False(t, result, "Returned value should be false")
}

func TestDownloadFileReturnsFalseOnError(t *testing.T) {
	var ircFileUrl = "irc://irc.org/#channel/botname/malfomed"
	var result = bridge.Download(ircFileUrl)
	assert.False(t, result, "Returned value should be false")

	ircFileUrl = "malformed"
	result = bridge.Download(ircFileUrl)
	assert.False(t, result, "Returned value should be false")
}
