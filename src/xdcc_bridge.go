package gazzettabot

type XdccBridge interface {
	Search(query string) []IrcFile
}

type IrcFile struct {
	name string
	url string
}
