package main

import (
	"github.com/KasiditR/websocket/controller"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	controller.SetupController(app)

	log.Fatal(app.Listen(":8080"))
}
