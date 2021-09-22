package main

import "fmt"

type hayvan interface {
	sayYourInfo()
}

func hayvanInterfaceFunction(hayvanlar ...hayvan) {
	for _, v := range hayvanlar {
		v.sayYourInfo()
	}
}

type kopek struct {
	isim, cins string
}

func (k kopek) sayYourInfo() {
	fmt.Println("Ben köpekgillerdenim,", "ismim:", k.isim, "cinsim:", k.cins)
}

type kedi struct {
	isim string
	boyu string
	hiz  string
}

func (k kedi) sayYourInfo() {
	fmt.Println("Ben kedigillerdenim,", "ismim:", k.isim, "boyum:", k.boyu, "hızım:", k.hiz)
}

type kemirgen struct {
	isim string
	kilo string
}

func (k kemirgen) sayYourInfo() {
	fmt.Println("Ben kemirgenim,", "ismim:", k.isim, "kilom:", k.kilo)
}

func main() {
	kopek := kopek{"Çomar", "Kangal"}
	kedi := kedi{"pamuk", "50cm", "10km"}
	fare := kemirgen{"Fare", "100gr"}

	hayvanInterfaceFunction(kopek, kedi, fare)
}
