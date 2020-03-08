package main

import (
  "fmt"
)

func pow(a []int)  {
  for i, v := range a {
    a[i] = v * v
  }
}

func main()  {
  // 参照渡し
  a := []int{1, 2, 3, 4, 5}
  pow(a)
  fmt.Println(a)
}
