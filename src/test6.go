package main

import (
  "fmt"
  "strconv"
)

func retStr(x, y int, str string) string {
  return strconv.Itoa(x + y) + str
}

func sumDiff(x, y int) (int, int) {
  sum := x + y
  diff := x - y
  return sum, diff
}

func multiDiv(x, y int) (multi, div int) {
  multi = x * y
  div = x / y
  return
}

func main()  {
  ret := retStr(1, 2, "hello")
  fmt.Println(ret)

  sum, diff := sumDiff(3, 4)
  fmt.Println(sum, diff)

  multi, div := multiDiv(6, 2)
  fmt.Println(multi, div)

  f := func (x, y int) int {
    return x + y
  }

  fmt.Printf("%#v\n", f)
  fmt.Printf("%#v\n", f(2,3))
}
