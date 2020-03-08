package main

import (
  "fmt"
)

func main()  {
  m := make(map[int]string)

  m[1] = "US"
  m[81] = "Japan"
  m[86] = "China"

  fmt.Println(m)

  names := map[int]string{
    1: "Taro",
    2: "Hanako",
    3: "Jiro",
  }

  fmt.Println(names)

  n := map[int][]int{
    1: {1, 2},
    2: {3, 4, 5},
  }

  fmt.Println(n)

  a := []map[string]int{
    {"col1": 1, "col2": 2, "col3": 3},
    {"col1": 4, "col2": 5, "col3": 6},
  }

  fmt.Println(a)
}
