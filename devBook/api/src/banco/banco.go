package banco

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Conectar() (*sql.DB, error) {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=devbook sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
