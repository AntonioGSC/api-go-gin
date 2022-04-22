package main

import (
	"github.com/AntonioGSC/api-go-gin/database"
	"github.com/AntonioGSC/api-go-gin/routes"
)

func main() {
	database.ConectaComDB()
	routes.HandleRequests()
}
