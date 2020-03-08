package main

import (
  "fmt"
)

func main()  {
  m := map[int]string{
    1: "foo",
    2: "bar",
    3: "baz",
  }

  for k, v := range m {
    fmt.Println(k, v)
  }

  delete(m, 1)

  fmt.Println(m)
}
