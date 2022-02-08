package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "testing.db"

func setupTest() *sql.DB {
	var db *sql.DB
	os.Remove(dbName)
	var err error
	db, err = sql.Open("sqlite3", "file:"+dbName+"?_fk=true") // fill out sql.Open function call with appropriate info
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")
	if err = InitDB(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB initialized")
	return db
}

func TestCreatePack(t *testing.T) {
	db := setupTest()
	defer db.Close()
	pack := Pack{Name: "RIP 12"}
	packdb, err := CreatePack(db, pack)
	if packdb.ID != 1 || packdb.Name != "RIP 12" || err != nil {
		t.Fatalf("CreatePack(db, pack) returned id = %v, name = %v, expected 1 and 'RIP 12'", packdb.ID, packdb.Name)
	}
}

/*
func main() {

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
*/
