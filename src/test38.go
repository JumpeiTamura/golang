package main

import (
  "fmt"
)

/*
type Stringer interface {
  String() string
}
*/

type mona struct {
  name string
  age int
}

func (m *mona) String() string {
  return `
  ∧＿∧　　／￣￣￣￣￣
（　´∀｀）＜　オマエモナー
（　　　　） 　＼＿＿＿＿＿
｜ ｜　|
（_＿）＿）`
}

func main()  {
  m := &mona{
    name: "foo",
    age: 13,
  }
  fmt.Println(m.name, m.age)
  fmt.Println(m)
}
