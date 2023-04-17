package main

import (
	"fmt"
)

func main() {
	fmt.Println("Найдем минимальное число из представленных вами!")

	fmt.Println("Введите первое число:")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Println("Введите второе число:")
	var secondNumber int
	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Println("Минимальное число - второе!")
	} else if firstNumber < secondNumber {
		fmt.Println("Минимальное число - первое!")
	} else if firstNumber == secondNumber {
		fmt.Println("Ваши числа равны!")
	}
}
