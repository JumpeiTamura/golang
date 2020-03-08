package main

import "fmt"

func main()  {
  a := [5]int{1, 2, 3, 4, 5}

  fmt.Println(a)

  b := [...]int{1,2,3,4,5,6}

  fmt.Println(b)

  b[5] = 7

  fmt.Println(b)
}
