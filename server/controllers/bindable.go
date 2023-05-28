package controllers

import "github.com/gofiber/fiber/v2"

// Bindable is a behaviour that is implemented by a REST API or a WebSocket handler,
// so it can be used with the router
type Bindable interface {
	Bind(app *fiber.App)
}
