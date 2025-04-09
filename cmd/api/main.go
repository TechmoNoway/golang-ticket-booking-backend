package main

import (
	"github.com/TechmoNoway/golang-ticket-booking-backend/handlers"
	"github.com/TechmoNoway/golang-ticket-booking-backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api/v1")

	handlers.NewEventHandler(server.Group("/events"), eventRepository)

	app.Listen(":4242")

}
