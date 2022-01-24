package main

import (
	"akabane_nyanko/backend/src/server"
)

func main() {
	router := server.GetRouter()
	router.Run()
}
