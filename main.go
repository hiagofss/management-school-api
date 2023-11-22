package main

import (
	"management-school/database"
	"management-school/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequests()
}
