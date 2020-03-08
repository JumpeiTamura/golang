package main

import (
  "fmt"
)

func main()  {
  const (
    A = 1
    B
    C
    D = "abc"
    E
  )
  const (
    X = iota
    Y
    Z
  )
  fmt.Println(A, B, C, D, E, X, Y, Z)
}
