package main

import (
	"github.com/gin-gonic/gin"
	"meeting-scheduler/config"
	"meeting-scheduler/models"
	"meeting-scheduler/routes"
)

func main() {
	r := gin.Default()
	db, err := config.InitDB() // Initialize the database connection
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.User{})

	// Set up routes
	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
