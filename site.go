package main
import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Site struct {
	ID int
	Name string
}

func (s Site) LoadDB() []Site {
	db, err := sql.Open("sqlite3", "./server.db")
	if err != nil {
		log.Printf("sql.Open not work::> %s", err)
	}
	rows, err := db.Query("SELECT * FROM Site")
	if err != nil {
		log.Printf("Select rows from Site has error ",err)
	}
	defer rows.Close()

	sites := []Site{}
	for rows.Next() {
		err = rows.Scan(&s.ID, &s.Name)
		if err != nil {
			log.Fatalf("rows.Scan Error %s", err)
		}
		sites = append(sites, s)
	}
	return sites
}
