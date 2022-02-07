package api

import (
	"database/sql"
)

func InitDB(db *sql.DB) error {
	query := `create table if not exists packs(
        id integer primary key autoincrement,
        name text not null,
        banner text
    );`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	query = `create table if not exists songs(
        id integer primary key autoincrement,
        title text not null,
        banner text,
        pack_fk integer,
        foreign key(pack_fk) references packs(id)
    );`
	_, err = db.Exec(query)

	return err
}

func CreateSong(db *sql.DB, song Song) (*Song, error) {
	res, err := db.Exec("insert into songs(title, banner, pack_fk) values (?,?,?)",
		song.Title, song.Banner, song.PackID)
	if err != nil {
		// handle error later
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	song.ID = id

	return &song, nil
}

func GetAllSongs(db *sql.DB) ([]Song, error) {
	rows, err := db.Query("select * from songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Banner, &song.PackID); err != nil {
			// handle error later
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func GetSong(db *sql.DB, id int64) (*Song, error) {
	var song Song
	err := db.QueryRow("select * from songs where id = ?", id).Scan(&song.ID, &song.Title, &song.Banner, &song.PackID)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func CreatePack(db *sql.DB, pack Pack) (*Pack, error) {
	res, err := db.Exec("insert into packs(name, banner) values (?,?)", pack.Name, pack.Banner)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	pack.ID = id

	return &pack, nil
}

func GetAllPacks(db *sql.DB) ([]Pack, error) {
	rows, err := db.Query("select * from packs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packs []Pack
	for rows.Next() {
		var pack Pack
		if err := rows.Scan(&pack.ID, &pack.Name, &pack.Banner); err != nil {
			// handle error later
			return nil, err
		}
		pack.Songs, _ = getSongsForPack(db, &pack)
		packs = append(packs, pack)
	}
	return packs, nil
}

func GetPack(db *sql.DB, id int64) (*Pack, error) {
	var pack Pack
	err := db.QueryRow("select * from packs where id = ?", id).Scan(&pack.ID, &pack.Name, &pack.Banner)
	if err != nil {
		return nil, err
	}
	pack.Songs, _ = getSongsForPack(db, &pack)
	return &pack, nil
}

func getSongsForPack(db *sql.DB, pack *Pack) ([]Song, error) {
	rows, err := db.Query("select * from songs where pack_fk = ?", pack.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Banner, &song.PackID); err != nil {
			// handle error later
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
