package main

import (
	"github.com/vishal2457/redmine-server/models"
	"github.com/vishal2457/redmine-server/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	router := routes.SetupRouter()

	// Initialize database
	models.SetupDatabase()

	// Start the HTTP API
	router.Run()
}
