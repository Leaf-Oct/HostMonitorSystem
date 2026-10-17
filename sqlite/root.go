package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var host_db *sql.DB

func InitDB(path string) error {
	host_db, err := sql.Open("sqlite", path)
	if err != nil {
		return err
	}
	host_db.SetMaxOpenConns(1)
	return nil
}

func CloseDB() {
	host_db.Close()
}
