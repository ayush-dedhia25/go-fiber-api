package main

import (
   "log"
   "github.com/gofiber/fiber/v2"
   "go-api/database"
)

func main() {
   // Create a new fiber instance
   app := fiber.New()
   // Try to connect to a mysql database
   database.ConnectDB()
   // Setup router
   SetupRoutes(app);
   // Start the server
   log.Fatal(app.Listen(":3000"))
}