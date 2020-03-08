package main

import (
  "fmt"
)

func main()  {
  L:
  for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
      k := i * j
      if k % 3 == 0 {
        continue L
      }
      fmt.Println(k)
    }
  }
}
