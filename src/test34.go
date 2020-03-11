package main

import (
  "fmt"
  "math"
)

type point struct {
  x float64
  y float64
}

func swap(p *point)  {
  p.x, p.y = p.y, p.x
}

func length(p *point) float64 {
  return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main()  {
  // new()で型へのポインタが得られる
  p := new(point)
  // 他の参照型と同様、要素にアクセスする際は参照する値に自動的に展開してくれるみたい
  p.x, p.y = 3.0, 4.0

  fmt.Println(p.x)

  swap(p)

  fmt.Println(p.x)

  fmt.Println(length(p))
}
