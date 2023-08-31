package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
	"todo/db"
	"todo/server"
)

func main() {
	database := db.Setup("postgres://hellofreshdev:hellofreshdev@localhost:5432/todo?sslmode=disable")
	router := server.Setup(database)

	println("Server starting!")

	_ = http.ListenAndServe(":8080", router)
}
