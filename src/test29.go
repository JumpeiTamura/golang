package main

import (
  "fmt"
)

func inc(i *int)  {
  *i++
}

func main()  {
  // 初めてのポインタ

  i := 5
  p := &i

  fmt.Println(*p)

  *p = 10
  fmt.Println(i)

  inc(&i)
  fmt.Println(i)
}
