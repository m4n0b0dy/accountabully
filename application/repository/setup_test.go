package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	path := "./testdata/test.db"
	_ = os.Remove(path)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	testRepo = NewSQLiteRepository(db)
	os.Exit(m.Run())
}
