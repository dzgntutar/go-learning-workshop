package main

import "fmt"

func main() {

	// Declaration
	fmt.Println("Declaration -->")

	/*
		var name, surname string = "Düzgün", "Tutar"
		language := "English"
		var age int = 34
		var isMarried bool = true

	*/

	/*
		var (
			name, surname string = "Düzgün", "Tutar"
			language             = "English"
			age           int    = 34
			isMarried     bool   = true
		)
	*/

	/*
		var name, surname, language, age, isMarried = "Düzgün", "Tutar", "English", 34, true
	*/

	name, surname, language, age, isMarried := "Düzgün", "Tutar", "English", 34, true

	fmt.Println(name)
	fmt.Println(surname)
	fmt.Println(language)
	fmt.Println(age)
	fmt.Println(isMarried)

	// Zero values
	fmt.Println("Zero values -->")

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
