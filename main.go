package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Print("__ Калькулятор индекса массы тела ___\n")
	fmt.Print("Введите ссввой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ссввой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(IMT, "\n")
}
