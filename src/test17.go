package main

import (
  "fmt"
)

func runDefer()  {
  defer fmt.Println(1)
  defer fmt.Println(2)
  defer fmt.Println(3)
  fmt.Println("done")
}

func runDefer2()  {
  defer runDefer()
  fmt.Println("end")
}

func runDefer3()  {
  defer func ()  {
    fmt.Println("a")
  }() // 関数呼び出し

  fmt.Println("finish")
}

func main()  {
  runDefer2()
  runDefer3()
}
