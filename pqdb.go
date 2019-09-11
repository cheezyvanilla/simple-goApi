package main

import(
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	connStr := "user=cheezyvanilla password=cheezy123 dbname=simpleAPI sslmode=verify-full"
	db,err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal(err)
	}
	return db
}