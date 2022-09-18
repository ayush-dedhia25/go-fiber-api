package utils

import (
   "time"
   "github.com/golang-jwt/jwt/v4"
   "go-api/config"
)

type MyApiClaim struct {
   Name     string   `json:"name"`
   Email    string   `json:"email"`
   jwt.RegisteredClaims
}

func CreateJwtToken(name, email string, uid uint) (string, error) {
   secretKey := config.GetEnv("SECRET_KEY")
   // Creating Claims/Payload for Jwt token
   claims := MyApiClaim{
      name,
      email,
      jwt.RegisteredClaims{
         ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
         Issuer: "Ayush Productions",
         ID: string(uid),
      },
   }
   // Getting a token object with all the standard token fields
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   // Generating Jwt Token and then return it
   tokenString, err := token.SignedString([]byte(secretKey))
   return tokenString, err
}

func ParseJwtToken(tokenString string) (*jwt.Token, error) {
   secretKey := config.GetEnv("SECRET_KEY")
   parser := jwt.NewParser()
   token, err := parser.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
      return []byte(secretKey), nil
   })
   return token, err
}