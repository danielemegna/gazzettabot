package gazzettabot

type XdccBridge interface {
	Search(query string) []IrcFile
}

type IrcFile struct {
	Name string
	SizeInMegaByte int
	Url string
}
