package main
import "database/sql"

type DB struct {
	*sql.DB
}

func NewDB(conn string) (*DB, error) {
	db, err := sql.Open("sqlite3", conn)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
