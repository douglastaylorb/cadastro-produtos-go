package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComDb() *sql.DB {
	conexao := "user=postgres dbname=registro_chamados password=douglastaylor#20 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
