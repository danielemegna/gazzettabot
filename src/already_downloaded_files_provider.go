package gazzettabot

type AlreadyDownloadedFilesProvider interface {
	List() []string
}
