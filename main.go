package main

import "fmt"

func userInput() float64 {
	var amount float64
	fmt.Scan(&amount)
	return amount
}

func convertCurrency(amount float64, srcCurrency string, dstCurrency string) {

}

func main() {
	const USD_EUR = 0.86
	const USD_RUB = 79.26
	const EUR_RUB = USD_RUB / USD_EUR
	var amount float64

	fmt.Print("Введите значение:")
	amount = userInput()
	fmt.Println(amount)
	fmt.Println(100 * EUR_RUB)
}
