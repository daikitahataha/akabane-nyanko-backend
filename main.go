package main

import (
	"akabane_nyanko/backend/src/db"
	"akabane_nyanko/backend/src/server"
)

func main() {
	router := server.GetRouter()

	db.Init()
	router.Run()
	db.Close()
}
