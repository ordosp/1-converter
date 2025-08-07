package main

import "fmt"

func main() {
	const USD_EUR = 0.86
	const USD_RUB = 79.26
	const EUR_RUB = USD_RUB / USD_EUR

	fmt.Println(100 * EUR_RUB)
}
