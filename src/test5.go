package main

import "fmt"

func main()  {
  var i interface{}

  fmt.Println(i)
  i = 1
  fmt.Println(i)
  i = false
  fmt.Println(i)
  i = "aiueo"
  fmt.Println(i)
  i = [...]int{1, 2, 3}
  fmt.Println(i)
}
