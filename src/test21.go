package main

import (
  "fmt"
)

func sum(a ...int) (n int) {
  for _, i := range a {
    n += i
  }
  return
}

func main()  {
  // スライスと可変長引数
  fmt.Println(sum())
  fmt.Println(sum(1, 2, 3))
  s := []int{4, 5, 6}
  fmt.Println(sum(s...))
}
