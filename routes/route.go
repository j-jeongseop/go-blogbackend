package routes

import (
    "github.com/gofiber/fiber/v3"
    "github.com/happynet78/go-blogbackend/controller"
)

func Setup(app *fiber.App) {
    app.Post("/api/register", controller.Register)
}
