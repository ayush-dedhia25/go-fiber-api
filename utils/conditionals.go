package utils

func IfThenElse(condition bool, a, b interface{}) interface{} {
   if condition {
      return a
   }
   return b
}