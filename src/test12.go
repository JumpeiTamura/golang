package main

import (
  "fmt"
)

func main()  {
  a := [...]string{
    "foo",
    "bar",
    "baz",
  }

  for i, s := range a {
    fmt.Println(i, s)
  }
}
