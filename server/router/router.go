package router

import (
	"salsa/config"
	"salsa/controllers"

	"github.com/gofiber/fiber/v2"
)

var server *fiber.App = nil

func Start() {
	err := server.Listen(":" + config.PortNumber())
	if err != nil {
		panic(err)
	}
}

func init() {
	server = fiber.New(fiber.Config{
		AppName: "Salsa",
	})

	for _, controller := range controllers.GetControllers() {
		controller.Bind(server)
	}
}
