package database

import (
   "fmt"
   "log"
   "gorm.io/gorm"
   "gorm.io/driver/sqlite"
   "go-api/models"
)

var DB *gorm.DB

func ConnectDB() {
   var err error

   // Try to connect with the Sqlite database
   DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
   if err != nil {
      log.Fatal("Failed to connect with database :(")
   }
   fmt.Println("Connection established successfully :)")
   
   // Auto pushing new changes to the database
   DB.AutoMigrate(&models.User{})
   fmt.Println("Database Migrated!")
}
