package main

import (
  "fmt"
)

func pow(p *[5]int)  {
  // *[...]intという指定はできないっぽい
  // 配列の参照渡しはスライスを使ったほうがよさそう
  for i, v := range p {
    p[i] = v * v
  }
}

func main()  {
  a := [...]int{1, 2, 3, 4, 5}

  pow(&a)

  fmt.Println(a)
}
