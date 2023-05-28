package apis

import "github.com/gofiber/fiber/v2"

type ExampleApi struct{}

func NewExampleApi() *ExampleApi {
	return &ExampleApi{}
}

func (e *ExampleApi) Bind(app *fiber.App) {
	greeterGroup := app.Group("/greet")
	greeterGroup.Get("/:name", e.handleGreet)
}

func (e *ExampleApi) handleGreet(ctx *fiber.Ctx) error {
	name := ctx.Params("name", "World")
	return ctx.SendString("Hello, " + name)
}
