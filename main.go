package main

import "fmt"

func inputAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите количество: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Это не число!")
		} else {
			return amount
		}
	}
}

func inputSrcCurrency() string {
	var srcCurrency string
	for {
		fmt.Print("Введите исходную валюту (USD EUR RUB): ")
		fmt.Scan(&srcCurrency)
		if srcCurrency == "USD" || srcCurrency == "EUR" || srcCurrency == "RUB" {
			return srcCurrency
		} else {
			fmt.Println("Введенное значение не является валютой. Возможные значения: USD EUR RUB")
		}
	}
}

func inputDstCurrency(srcCurrency string) string {
	var dstCurrency string

	// формируем список возможных значений целевой валюты
	var currencies string
	if srcCurrency != "USD" {
		currencies = currencies + "USD "
	}
	if srcCurrency != "EUR" {
		currencies = currencies + "EUR "
	}
	if srcCurrency != "RUB" {
		currencies = currencies + "RUB "
	}

	for {
		fmt.Printf("Введите целевую валюту (%s): ", currencies)
		fmt.Scan(&dstCurrency)
		if srcCurrency != "" && dstCurrency == srcCurrency {
			fmt.Println("Эта валюта не может быть целевой, так как она указана как исходная!")
			continue
		}
		if dstCurrency == "USD" || dstCurrency == "EUR" || dstCurrency == "RUB" {
			return dstCurrency
		} else {
			fmt.Printf("Введенное значение не является валютой. Возможные значения: %s\n", currencies)
		}
	}
}

func convertCurrency(amount float64, srcCurrency string, dstCurrency string) float64 {
	var result float64

	const USD_EUR = 0.86
	const USD_RUB = 79.26
	const EUR_RUB = USD_RUB / USD_EUR

	fmt.Println(amount, srcCurrency, dstCurrency)

	switch {
	case srcCurrency == "USD" && dstCurrency == "EUR":
		result = amount * USD_EUR
	case srcCurrency == "USD" && dstCurrency == "RUB":
		result = amount * USD_RUB
	case srcCurrency == "EUR" && dstCurrency == "USD":
		result = amount / USD_EUR
	case srcCurrency == "EUR" && dstCurrency == "RUB":
		result = amount * EUR_RUB
	case srcCurrency == "RUB" && dstCurrency == "EUR":
		result = amount / EUR_RUB
	case srcCurrency == "RUB" && dstCurrency == "USD":
		result = amount / USD_RUB
	}

	return result
}

func main() {
	var srcCurrency string
	var dstCurrency string
	var amount float64
	var result float64

	srcCurrency = inputSrcCurrency()
	amount = inputAmount()
	dstCurrency = inputDstCurrency(srcCurrency)

	result = convertCurrency(amount, srcCurrency, dstCurrency)
	fmt.Println(result)
}
