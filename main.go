package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ivanes285/GoPostgresAPI/routes/v1"
)

func main() {
	app := fiber.New()

	//Routes
	app.Get("/", routes.Home)


	app.Listen(":3000")
}
