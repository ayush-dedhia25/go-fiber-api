package models

import (
   "fmt"
   "errors"
   "gorm.io/gorm"
   "golang.org/x/crypto/bcrypt"
)

type User struct {
   gorm.Model
   Name        string   `json:"name"`
   Email       string   `json:"email" gorm:"unique"`
   Password    string   `json:"password"`
   Age         uint8    `json:"age"`
}

func (u *User) BeforeCreate(txn *gorm.DB) error {
   hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
   if err != nil {
      fmt.Println(err)
      return err
   }
   u.Password = string(hashed)
   return nil
}

func (u *User) IsValidPass(plainTextPass string) bool {
   err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainTextPass))
   if err != nil {
      fmt.Println(err)
      return false
   }
   return true
}

func (u *User) Validate() error {
   if u.Name == "" || u.Name == " " {
      return errors.New("Name is or not provided!")
   }
   if u.Email == "" || u.Email == " " {
      return errors.New("Email is invalid or not provided!")
   }
   if u.Password == "" || u.Password == " " {
      return errors.New("Password is not provided!")
   }
   if u.Age < 18 || u.Age > 99 {
      return errors.New("Age is not allowed!")
   }
   return nil
}