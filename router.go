package main

import (
   "github.com/gofiber/fiber/v2"
   "github.com/gofiber/fiber/v2/middleware/logger"
   "go-api/routes"
)

func SetupRoutes(app *fiber.App) {
   api := app.Group("/api", logger.New())
   routes.SetupUserRoutes(api)
   routes.SetupAuthRoutes(api)
}