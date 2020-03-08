package main

import (
  "fmt"
)

func integers(x int) func() int {
  i := x
  return func () int {
    i += 1
    return i
  }
}

func main()  {
  ints := integers(5)
  fmt.Println(ints())
  fmt.Println(ints())
  fmt.Println(ints())
}
