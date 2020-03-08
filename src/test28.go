package main

import (
  "fmt"
)

func main()  {
  ch1 := make(chan int, 1)
  ch2 := make(chan int, 1)
  ch3 := make(chan int, 1)
  ch1 <- 1
  ch2 <- 2

  var i int
  select {
  case i = <- ch1:
    fmt.Println(i)
  case i = <- ch2:
    fmt.Println(i)
  case ch3 <- 3:
    i = <- ch3
    fmt.Println(i)
  }
}
