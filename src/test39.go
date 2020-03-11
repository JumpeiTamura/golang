package main

import (
  "fmt"
)

type callable interface {
  call() string
}

type phone struct {
  number string
}

type spilit struct {
  owner string
  wish string
}

func (p *phone) call() string {
  return fmt.Sprintf("Calling: %s ...", p.number)
}

func (s *spilit) call() string {
  return fmt.Sprintf("Your wish '%s' will be the truth!", s.wish)
}

func main()  {
  // 共通のインターフェースを使う（ダックタイピング）

  p := &phone{
    number: "090-1234-5678",
  }

  s := &spilit{
    owner: "aladin",
    wish: "genie, you are free!",
  }

  c := []callable{
    p,
    s,
  }

  for _, i := range c {
    fmt.Println(i.call())
  }
}
