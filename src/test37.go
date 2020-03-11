package main

import (
  "fmt"
)

/*
type error interface {
  Error() string
}
*/

type myError struct {
  message string
  errcode int
}

// error型のインターフェースError()を実装
func (e *myError) Error() string {
  return e.message
}

func raiseError() error {
  // インターフェースは参照型なので型のアドレスを返す
  return &myError{
    message: "エラーが発生しました",
    errcode: 1234,
  }
}

func main()  {
  err := raiseError()
  fmt.Println(err.Error())

  // 型アサーションで本来の型を取り戻す
  if e, ok := err.(*myError); ok {
    fmt.Println(e.errcode)
  }
}
