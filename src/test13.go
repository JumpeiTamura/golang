package main

import (
  "fmt"
)

func main()  {
  a := 5

  switch a {
  case 1, 2:
    fmt.Println("1 or 2")
  case 3, 4, 5:
    fmt.Println("3 or 4 or 5")
  }
}
