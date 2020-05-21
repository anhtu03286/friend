package main

import (
	"github.com/anhtu03286/friend/config"
	"github.com/anhtu03286/friend/router"
)

func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	router.HandleRequest(db)
}