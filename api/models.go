package api

// TODO: add song artist and step artist
type Song struct {
	ID       int64
	Title    string
	Subtitle string
	Artist   string
	Banner   string
	PackID   int64
}

type Pack struct {
	ID     int64
	Name   string
	Banner string
	Songs  []Song
}
