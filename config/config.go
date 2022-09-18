package config

import (
   "os"
   "fmt"
   "github.com/joho/godotenv"
)

func GetEnv(key string) string {
   err := godotenv.Load(".env")
   if err != nil {
      fmt.Print("Error loading .env")
   }
   return os.Getenv(key)
}