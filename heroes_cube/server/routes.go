package server

import "github.com/gofiber/fiber/v3"

func HealthCheckHandler(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "message": "Heroes Cube API is running!"})
}
