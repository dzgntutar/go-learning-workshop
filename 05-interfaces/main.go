package main

import (
	"fmt"
	"math"
)

//**************************************
type sekil interface {
	alan() float64
	cevre() float64
}

//**************************************
type dikdortgen struct {
	a int
	b int
}

func (d dikdortgen) alan() float64 {
	return (float64)(d.a * d.b)
}
func (d dikdortgen) cevre() float64 {
	return float64((d.a + d.b) * 2)
}

//**************************************
type daire struct {
	yaricap float64
}

func (d daire) alan() float64 {
	return math.Pi * d.yaricap * d.yaricap
}
func (d daire) cevre() float64 {
	return 2 * math.Pi * d.yaricap
}

func interfacegoster(s sekil) {
	fmt.Println("Alan: ", s.alan())
	fmt.Println("Çevre: ", s.cevre())
}

func main() {
	var dikdortgen1 dikdortgen
	dikdortgen1.a = 15
	dikdortgen1.b = 60

	interfacegoster(dikdortgen1)

	fmt.Println("****************************************")
	var daire1 daire
	daire1.yaricap = 8
	interfacegoster(daire1)

}
