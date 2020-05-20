package main

import (
	"github.com/anhtu03286/friend/config"
	"github.com/anhtu03286/friend/router"
)

// @title Friend Management API
// @version 1.0
// @description This is a sample service for managing Friend Management
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	db, _ := config.OpenDB()
	defer db.Close()

	router.HandleRequest(db)
}