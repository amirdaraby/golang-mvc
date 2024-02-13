package routes

import (
	"github.com/amirdaraby/golang-mvc/controllers"
	"github.com/gofiber/fiber/v2"
)

const apiPrefix = "/api"

func Register(app *fiber.App) {
	// use api to add new routes
	api := app.Group(apiPrefix)

	api.Post("user", controllers.StoreUser)
}
