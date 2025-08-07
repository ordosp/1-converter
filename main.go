package main

import "fmt"

func main() {
  const USD_EUR = 0.86
  const USD_RUB = 79.26

  fmt.Println(100 / USD_EUR * USD_RUB)
}
