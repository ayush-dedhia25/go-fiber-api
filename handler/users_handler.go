package handler

import (
   "errors"
   "fmt"
   "github.com/gofiber/fiber/v2"
   "gorm.io/gorm"
   "go-api/database"
   "go-api/models"
)

func GetUsers(c *fiber.Ctx) error {
   db := database.DB
   var users []models.User
   db.Find(&users)
   return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
   db := database.DB
   // Getting id from route params
   userId := c.Params("id")
   
   var user models.User
   // Checking if the user exists or not in the database
   err := db.First(&user, userId).Error
   if errors.Is(err, gorm.ErrRecordNotFound) {
      return c.Status(404).JSON(fiber.Map{
         "message": fmt.Sprintf("No user found with ID = %s", userId),
      })
   }
   
   return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
   // Getting database connection
   db := database.DB
   user := new(models.User)
   
   // Parsing request body to user struct
   if err := c.BodyParser(user); err != nil {
      return c.JSON(fiber.Map{"message": err.Error()})
   }
   
   // Checking if the user already exists in the database or not
   // in this case user 'should not' already exist to create a fresh account
   var existingUser models.User
   err := db.First(&existingUser, "email = ?", user.Email).Error
   // This tells that User was not found i.e. its safe to proceed
   if !errors.Is(err, gorm.ErrRecordNotFound) {
      return c.Status(409).JSON(fiber.Map{
         "error": "Email is already in use...!",
      })
   }
   
   // Saving user in the database
   if err := db.Create(user).Error; err != nil {
      return c.Status(500).JSON(fiber.Map{
         "message": "Oops!! Something went wrong!",
      })
   }
   
   return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
   db := database.DB
   // Getting id from route params
   userId := c.Params("id")
   
   // Checking for the user's existence
   var user models.User
   err := db.First(&user, userId).Error
   if errors.Is(err, gorm.ErrRecordNotFound) {
      return c.Status(404).JSON(fiber.Map{
         "message": fmt.Sprintf("No user found with ID = %s", userId),
      })
   }
   
   // Creating a patch object for the update
   var updateUserData models.UpdateUser
   if err = c.BodyParser(&updateUserData); err != nil {
      return c.JSON(fiber.Map{"message": err.Error()})
   }
   
   // Updating the user object and saving it in the database 
   if updateUserData.Name != "" {
      user.Name = updateUserData.Name
   }
   if updateUserData.Age != 0 {
      user.Age = updateUserData.Age
   }
   // Finally updating the user in the database 
   db.Save(&user)
   
   return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
   db := database.DB
   // Getting id from route params
   userId := c.Params("id")
   
   var user models.User
   err := db.First(&user, userId).Error
   if errors.Is(err, gorm.ErrRecordNotFound) {
      return c.Status(404).JSON(fiber.Map{
         "message": fmt.Sprintf("No user found with ID = %s", userId),
      })
   }
   
   err = db.Delete(&user, userId).Error
   if err != nil {
      return c.Status(500).JSON(fiber.Map{
         "error": "Failed to delete user!",
      })
   }
   
   return c.JSON(fiber.Map{"message": "User deleted!"})
}