package main

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", Register)
	app.Post("/api/login", Login)
	app.Get("/api/user", UserSession)
	app.Post("/api/logout", Logout)
}