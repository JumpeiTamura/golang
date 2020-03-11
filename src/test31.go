package main

import (
  "fmt"
)

type Callback func(int) int

func sum(ints []int, callback Callback) (sum int) {
  for _, i := range ints {
    sum += callback(i)
  }
  return
}

func main()  {
  a := []int{1, 2, 3}

  n := sum(
    a,
    func (i int) int {
      return i * i
    },
  )

  fmt.Println(n)
}
