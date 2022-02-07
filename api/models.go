package api

type Song struct {
	ID     int64
	Title  string
	Banner string
	PackID int64
}

type Pack struct {
	ID     int64
	Name   string
	Banner string
	Songs  []Song
}
