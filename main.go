package main

import (
	"api_rest_gin/database"
	"api_rest_gin/routes"
)

func main() {
	database.ConnectPostgres()
	routes.Handle()
}
