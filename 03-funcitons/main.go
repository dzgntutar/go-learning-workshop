package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumInt(25, 50))
	fmt.Println(SumFloat(12.75, 50.48))

	bolum, kalan := bolme(98, 4)
	fmt.Println(bolum)
	fmt.Println(kalan)
}

func SumInt(number1, number2 int) int {
	return number1 + number2
}

func SumFloat(number1, number2 float64) float64 {
	return number1 + number2
}

func bolme(bolunen, bolen int) (bolum, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum, kalan
}
