package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBancoDeDados() *sql.DB {
	conexao := "user=root dbname=root password=root host=0.0.0.0 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
