package routes

import (
   "github.com/gofiber/fiber/v2"
   "go-api/handler"
)

func SetupAuthRoutes(router fiber.Router) {
   auth := router.Group("/auth")
   auth.Post("/login", handler.SignIn)
   auth.Post("/", handler.CreateUser)
}