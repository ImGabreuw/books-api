package main

import "books-api/server"

func main() {
	s := server.NewServer()

	s.Run()
}
