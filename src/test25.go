package main

import (
  "fmt"
)

func main()  {
  m := map[int]int{
    1: 2,
    3: 4,
    5: 6,
  }

  fmt.Println(m)

  a, ok := m[1]
  fmt.Println(a, ok)

  // okは再定義しても問題ない(bが新しい変数だから)
  // https://qiita.com/Jxck_/items/2616abafea89ee97c477
  b, ok := m[2]
  fmt.Println(b, ok)

  // b, ok := m[3] // 同じ変数だけだとダメ
  // 同じ理屈で、以下もOK
  b, ng := m[3]
  fmt.Println(b, ng)

  if c, ok := m[5]; ok {
    fmt.Println("値は", c)
  }else{
    fmt.Println("そんなもんはいないぞい")
  }
}
