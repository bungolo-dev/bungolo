package infrustucture

import (
	"database/sql"

	_ "embed"

	"github.com/bungolow-dev/bungolow/pkg/application/repository"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() repository.DbConnection {
	return &Database{
		path: "./bungolow.db",
	}
}

type Database struct {
	path string
}

func (database *Database) Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", database.path)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (database *Database) Initialize() error {
	db, err := sql.Open("sqlite3", "./bungolow.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS rooms (
		id STRING PRIMARY KEY NOTNULL,
		name TEXT NOT NULL,
		description TEXT
		image TEXT
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
