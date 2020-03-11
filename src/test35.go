package main

import (
  "fmt"
  "math"
)

type point struct {
  x float64
  y float64
}

// メソッド定義
func (p *point) length() float64 {
  return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main()  {
  p := new(point)
  p.x, p.y = 3, 4

  fmt.Println(p.length())
}
