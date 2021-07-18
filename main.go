package main

import (
	"books-api/database"
	"books-api/server"
)

func main() {
	database.StartDB()

	s := server.NewServer()

	s.Run()
}
