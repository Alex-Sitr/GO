package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела ___")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
}

func calculateIMT(usKg float64, usHeight float64) float64 {
	const IMTPower = 2
	IMT := usKg / math.Pow(usHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
