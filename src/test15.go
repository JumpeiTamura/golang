package main

import (
  "fmt"
)

func main()  {
  var x interface{} = 3.14

  switch v := x.(type) {
  case int:
    fmt.Println(v, "is int")
  case float64:
    fmt.Println(v, "is float")
  case string:
    fmt.Println(v, "is string")
  default:
    fmt.Println("unknown")
  }
}
