package main

import (
  "fmt"
)

func main()  {
  // スライスの落とし穴
  a := [5]int{1, 2, 3, 4, 5}

  s := a[1:3]
  fmt.Println(a)
  fmt.Println(s)

  //　配列の変更はスライスにも反映される
  a[1] = 1
  fmt.Println(a)
  fmt.Println(s)

  // スライスの変更は配列にも反映される
  s[1] = 2
  fmt.Println(a)
  fmt.Println(s)

  // キャパ内でスライスをappendすると元の配列も上書きされる
  fmt.Println(cap(s))
  s = append(s, 1)
  fmt.Println(a)
  fmt.Println(s)
  s[0] = 0
  fmt.Println(a)
  fmt.Println(s)

  // キャパオーバーすると元の配列とは別の領域を参照するようになる
  s = append(s, 2, 3)
  fmt.Println(a)
  fmt.Println(s)
}
