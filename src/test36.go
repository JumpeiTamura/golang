package main

import (
  "fmt"
  "math"
)

type point struct {
  x int
  y int
  length float64
}

func newPoint(x, y int) *point {
  // p := &point{x: x, y: y}
  p := new(point)
  p.x, p.y = x, y
  p.length = math.Sqrt(float64(x*x + y*y))
  return p
}

func main()  {
  p := newPoint(3, 4)

  fmt.Println(*p)
}
