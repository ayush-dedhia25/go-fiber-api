package routes

import (
   "github.com/gofiber/fiber/v2"
   mid "go-api/middlewares"
   "go-api/handler"
)

func SetupUserRoutes(router fiber.Router) {
   users := router.Group("/users", mid.AuthMiddleware)
   
   users.Get("/", handler.GetUsers)
   users.Get("/:id", handler.GetUser)
   users.Put("/:id", handler.UpdateUser)
   users.Delete("/:id", handler.DeleteUser)
}