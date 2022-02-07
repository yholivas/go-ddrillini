package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yholivas/ddrillini-go-api/api"
)

var db *sql.DB

const dbName = "ddrillini.db"

func main() {
	os.Remove(dbName)
	var err error
	db, err = sql.Open("sqlite3", "file:"+dbName+"?_fk=true") // fill out sql.Open function call with appropriate info
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	if err = api.InitDB(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB initialized")

	var packdb *api.Pack
	pack := api.Pack{Name: "RIP 12"}
	packdb, err = api.CreatePack(db, pack)
	if err != nil {
		log.Fatal(err)
	}

	var song1db, song2db *api.Song
	song1 := api.Song{Title: "9mm", PackID: packdb.ID}
	song2 := api.Song{Title: "Earthquake", PackID: packdb.ID}
	song1db, err = api.CreateSong(db, song1)
	if err != nil {
		log.Fatal(err)
	}
	song2db, err = api.CreateSong(db, song2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(song1db)
	fmt.Println(song2db)
	fmt.Println(api.GetPack(db, song1db.ID))
}
