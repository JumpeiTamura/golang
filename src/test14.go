package main

import (
  "fmt"
)

func main()  {
  // 型アサーション
  var a interface{} = 3.14

  i, isInt := a.(int)
  fmt.Println(i, isInt)

  // goにはfloatという型はないっぽい
  f, isFroat := a.(float64)
  fmt.Println(f, isFroat)

  s, isString := a.(string)
  fmt.Println(s, isString)
}
