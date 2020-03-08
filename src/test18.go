package main

import (
  f "fmt"
  r "runtime"
)

func main()  {
  f.Println("NumCPU:", r.NumCPU())
  f.Println("NumGoroutine:", r.NumGoroutine())
  f.Println("Version:", r.Version())
}
