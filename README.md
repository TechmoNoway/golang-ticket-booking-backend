# Golang Ticket Booking Backend

A backend API service for ticket booking management built with Go and Fiber framework.

## Project Overview

This project implements a RESTful API for managing events and ticket bookings. It follows a clean architecture pattern with separate layers for handlers, models, and repositories.

## Technologies

- Go 1.24.2
- [Fiber](https://github.com/gofiber/fiber) - Web framework
- Clean Architecture Pattern

## Project Structure

```
├── cmd/                  # Command line applications
│   └── api/              # Main API entry point
├── handlers/             # HTTP request handlers
├── models/               # Domain models and interfaces
├── repositories/         # Data access implementations
└── bin/                  # Compiled binaries
```

## API Endpoints

- `GET /api/v1/events` - Retrieve all events
- `GET /api/v1/events/:eventId` - Retrieve a specific event
- `POST /api/v1/events` - Create a new event

## Getting Started

### Prerequisites

- Go 1.24.2 or higher

### Installation

1. Clone the repository
```bash
git clone https://github.com/TechmoNoway/golang-ticket-booking-backend.git
cd golang-ticket-booking-backend
```

2. Install dependencies
```bash
go mod download
```

### Running the Application

Use the included Makefile to build and start the application:

```bash
make start
```

This will compile the application and start the server on port 4242.

### Building the Application

```bash
make build
```

## License

MIT