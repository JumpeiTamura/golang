package main

import (
  "fmt"
)

func main()  {
  // 構造体を使った共通化
  // オブジェクト指向の継承みたいなもの？

  type base struct {
    id int
    name string
  }

  type user struct {
    base
    email string
  }

  type country struct {
    base
    capital string
  }

  foo := user{
    base: base{
      // base.id: 1 みたいな書き方はできないみたい
      id: 1,
      name: "foo",
    },
    email: "foo@barbaz.com",
  }

  japan := country{
    base: base{
      id: 1,
      name: "Japan",
    },
    capital: "Tokyo",
  }

  fmt.Println(foo.base.name)
  fmt.Println(japan.name)

  // 読み書き両方OK
  foo.name = "Hanako"

  fmt.Println(foo.name)
}
