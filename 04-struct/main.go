package main

import "fmt"

func main() {
	var student struct {
		name string
		age  int
	}

	fmt.Println(student)

	student.name = "Düzgün Tutar"
	student.age = 34

	fmt.Println(student.name)
	fmt.Println(student.age)

	fmt.Println("**********************************")

	type teacher struct {
		name string
		age  int
		kids []string
	}

	var teacher1 teacher
	teacher1.name = "Arda"
	teacher1.age = 25

	var teacher2 teacher
	teacher2.name = "Kaan"
	teacher2.age = 28

	fmt.Println(teacher1)
	fmt.Println(teacher2)

	fmt.Println("**********************************")

	teacher3 := teacher{
		name: "Musa",
		age:  36,
		kids: []string{
			"Melis",
			"Can",
		},
	}

	fmt.Printf("%#v\n", teacher3)

}
