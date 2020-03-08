package main

import (
  "fmt"
  "time"
)

func receiver(ch <-chan int)  {
  for {
    // 受信
    if i, ok := <- ch; ok{
      fmt.Println(i)
    }else{
      break
    }
  }
}

func main()  {
  // 送受信ができるチャネルじゃないとダメ
  ch := make(chan int, 20)

  go receiver(ch)

  for i := 0; i < 10; i++ {
    // 送信
    ch <- i
  }
  close(ch)

  // mainの処理が終わるとプログラム自体終わっちゃうみたいなので、待つ
  time.Sleep(1 * time.Second)
}
