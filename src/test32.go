package main

import (
  "fmt"
)

type point struct {
  x, y int
}

func swap(p *point)  {
  p.x, p.y = p.y, p.x
}

func main()  {
  // 構造体のポインタ
  // Goでは頻出らしい
  p := &point{x:3, y:4}

  fmt.Println(*p)

  swap(p)

  fmt.Println(*p)
}
