package handler

import (
   "errors"
   //"fmt"
   "github.com/gofiber/fiber/v2"
   "gorm.io/gorm"
   "go-api/database"
   "go-api/models"
   "go-api/utils"
)

func SignIn(c *fiber.Ctx) error {
   db := database.DB
   
   loginData := new(models.Login)
   if err := c.BodyParser(loginData); err != nil {
      return c.Status(400).JSON(fiber.Map{"error": err.Error()})
   }
   
   if err := loginData.Validate(); err != nil {
      return c.Status(400).JSON(fiber.Map{"error": err.Error()})
   }
   
   var user models.User
   err := db.First(&user, "email = ?", loginData.Email).Error
   if errors.Is(err, gorm.ErrRecordNotFound) {
      return c.Status(404).JSON(fiber.Map{
         "error": "Please check your credentials...!",
      })
   }
   
   if !user.IsValidPass(loginData.Password) {
      return c.Status(400).JSON(fiber.Map{
         "error": "Please check your email or password!",
      })
   }
   
   token, err := utils.CreateJwtToken(user.Name, user.Email, user.ID)
   if err != nil {
      return c.Status(500).JSON(fiber.Map{"error": "Oops! Something went wrong!"})
   }
   
   return c.Status(200).JSON(fiber.Map{"access_token": token})
}