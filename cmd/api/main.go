package main

import (
	"github.com/TechmoNoway/golang-ticket-booking-backend/config"
	"github.com/TechmoNoway/golang-ticket-booking-backend/db"
	"github.com/TechmoNoway/golang-ticket-booking-backend/handlers"
	"github.com/TechmoNoway/golang-ticket-booking-backend/middlewares"
	"github.com/TechmoNoway/golang-ticket-booking-backend/repositories"
	"github.com/TechmoNoway/golang-ticket-booking-backend/services"
	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api/v1")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(privateRoutes.Group("/events"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/tickets"), ticketRepository)

	app.Listen(":" + envConfig.ServerPort)

}
