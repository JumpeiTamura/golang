package main

import (
  "fmt"
  r "reflect"
)

func main()  {
  s := make([]int, 5, 10)
  fmt.Println(s, len(s), cap(s))

  a := []int{1, 2, 3}
  fmt.Println(a, len(a), cap(a))

  b := [...]int{1, 2, 3, 4, 5}
  fmt.Println(b, r.TypeOf(b))

  c := b[0:3]
  fmt.Println(c, r.TypeOf(c))

  d := b[:]
  d = append(d, 6)
  fmt.Println(d, r.TypeOf(d))
}
