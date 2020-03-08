package main

import (
  "fmt"
)

func even(x int) (int, bool) {
  b := false
  if x % 2 == 0 {
    b = true
  }
  return (x/2), b
}

func main()  {
  if x, err := even(5); err == false {
    fmt.Println("Error!")
  }else{
    fmt.Println(x)
  }
}
