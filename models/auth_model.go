package models

import "errors"

type UpdateUser struct {
   Name string `json:"name" validate:"required"`
   Age  uint8  `json:"age" validate:"required"`
}

type Login struct {
   Email    string `json:"email" validate:"required,email"`
   Password string `json:"password" validate:"required"`
}

func (up *UpdateUser) Validate() error {
   if up.Name == "" || up.Name == " " {
      return errors.New("Name is not provided!")
   }
   if up.Age < 18 || up.Age > 99 {
      return errors.New("Age is not valid!")
   }
   return nil
}

func (l *Login) Validate() error {
   if l.Email == "" || l.Email == " " {
      return errors.New("Email is not provided!")
   }
   if l.Password == "" || l.Password == " " {
      return errors.New("Password is not provided!")
   }
   return nil
}