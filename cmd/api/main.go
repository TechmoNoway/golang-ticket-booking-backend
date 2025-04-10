package main

import (
	"github.com/TechmoNoway/golang-ticket-booking-backend/config"
	"github.com/TechmoNoway/golang-ticket-booking-backend/db"
	"github.com/TechmoNoway/golang-ticket-booking-backend/handlers"
	"github.com/TechmoNoway/golang-ticket-booking-backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api/v1")

	handlers.NewEventHandler(server.Group("/events"), eventRepository)

	app.Listen(":" + envConfig.ServerPort)

}
