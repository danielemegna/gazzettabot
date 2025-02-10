package gazzettabot

import (
	"github.com/samber/lo"
	"log"
	"os"
)

type FileSystemAlreadyDownloadedFilesProvider struct {
	DownloadFolderPath string
}

func (this FileSystemAlreadyDownloadedFilesProvider) List() []string {
	var entries, err = os.ReadDir(this.DownloadFolderPath)
	if err != nil {
		log.Fatal("Error reading download folder! - ", err)
	}
	return lo.Map(entries, func(e os.DirEntry, _ int) string { return e.Name() })
}
