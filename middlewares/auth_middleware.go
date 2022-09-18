package middlewares

import (
   "strings"
   "github.com/gofiber/fiber/v2"
   "github.com/golang-jwt/jwt/v4"
   "go-api/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
   // Extract the 'authorization' header from request headers
   rawAuthHeader := c.Request().Header.Peek("authorization")
   
   // By default the header is a bytearray so...
   // Firstly converting it to the string.
   // Secondly checking if the headers value starts with the `Bearer` prefix.
   if strings.HasPrefix(string(rawAuthHeader), "Bearer ") {
      // Split the header's value to get the `Bearer` & actual token.
      tags := strings.Split(string(rawAuthHeader), "Bearer ")
      
      // Check for the token's existence in the header's value
      if tags[1] != "" || tags[1] != " " {
         // Parse it to check the token's validation
         token, err := utils.ParseJwtToken(tags[1])
         if token.Valid {
            // Token is valid here...
            return c.Next()
         } else if ve, ok := err.(*jwt.ValidationError); ok {
            // Here casting/converting the err that we received above to the jwt's ValidationError
            // To handle the appropriate token errors.
            if ve.Errors & jwt.ValidationErrorMalformed != 0 {
               // Token is not a token
               return c.Status(403).JSON(fiber.Map{"error": "Unauthorised access"})
            } else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
               // Token is either expired or not yet valid
               // Not yet valid here means token is probably used before its validity or a.k.a (Not Before Time) claim.
               return c.Status(403).JSON(fiber.Map{
                  "error": "Session has been expired please try to re-login and come back",
               })
            } else {
               return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
            }
         } else {
            return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
         }
      }
   }
   // Here probably Unauthorization header is missing
   return c.Status(500).JSON(fiber.Map{"error": "Unauthorised Access"})
}