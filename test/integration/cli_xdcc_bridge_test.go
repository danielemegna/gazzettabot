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
	var ircFile = IrcFile{
		Name: "Star.Wars.Rebels.3x20.Ultimo.Atto.Prima.Parte.ITA.ENG.DLMux.XviD-Pir8.srt",
		Url:  "irc://irc.irc-files.org/#FaNtAsYlAnD/FaNtAsYlAnD|AnImE|01/3112",
	}

	bridge.DownloadOneOf([]IrcFile{ircFile})

	assertAndDeleteDonwloadedFile(t, ircFile, 175)
}

func TestDownloadNextFileOnUnreachabileFile(t *testing.T) {
	t.Skip("Very slow test, it use the timeout - todo move it in env var")
	var validIrcFile = IrcFile{
		Name: "Naruto.Naruto.107.-.Sasuke.contro.Naruto.srt",
		Url:  "irc://irc.irc-files.org/#FaNtAsYlAnD/FaNtAsYlAnD|AnImE|01/4919",
	}
	var ircFiles = []IrcFile{
		{
			Name: "Malformed.url",
			Url:  "malformed",
		},
		{
			Name: "Malformed.irc.url",
			Url:  "irc://irc.org/#channel/botname/malfomed",
		},
		{
			Name: "Unreachabile.File.pdf",
			Url:  "irc://irc.openjoke.org/#TILT/TLT|EDICOLA|01/1190",
		},
		validIrcFile,
	}

	bridge.DownloadOneOf(ircFiles)

	assertAndDeleteDonwloadedFile(t, validIrcFile, 16045)
}

func assertAndDeleteDonwloadedFile(t *testing.T, ircFile IrcFile, expectedSizeInBytes int) {
	var downloadFolder = os.Getenv("DOWNLOAD_FOLDER")
	var expectedDownloadedFilepath = downloadFolder + "/" + ircFile.Name
	var fileInfo, err = os.Stat(expectedDownloadedFilepath)
	assert.Nil(t, err, "Cannot find expected downloaded file")
	assert.NotNil(t, fileInfo, "Cannot find expected downloaded file")
	assert.Equal(t, int64(expectedSizeInBytes), fileInfo.Size())
	err = os.Remove(downloadFolder + "/" + ircFile.Name)
	assert.Nil(t, err, "Cannot delete downloaded file")
}
