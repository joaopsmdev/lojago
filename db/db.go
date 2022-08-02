package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	//conexao com o banco postgress
	conexao := "user=postgres dbname=LojaDoRafinha password=pgadmin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
