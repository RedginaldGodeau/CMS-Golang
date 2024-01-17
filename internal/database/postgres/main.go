package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Connection() (*sql.DB, error) {
	var db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		db.Close()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
