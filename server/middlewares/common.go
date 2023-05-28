package middlewares

import (
	"salsa/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// AllowJSON sets content type of both the request and response to application/json
func AllowJSON(c *fiber.Ctx) error {
	c.Request().Header.Set("Content-Type", "application/json")
	c.Set("Content-Type", "application/json")
	return c.Next()
}

func AllowMethods(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	return c.Next()
}

func AllowCors(c *fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowOrigins: config.AllowedClients(),
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	})(c)
}
func AllowHeaders(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Headers", "Origin,Content-Type,Accept")
	return c.Next()
}
