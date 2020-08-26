package main

import "fmt"

// main function declaration
func main(){
	// studentsA is a variable or type slice
	studentsA := [] string { "Sanket", "Rohit"}

	// printing of list
	fmt.Println("Students of section A : ")
	for i, student := range studentsA {
		fmt.Println(i,student)
}

	// studentsB assign a custom type variable
	studentsB := classroom{"Swati","Gargi"}

	// calling a function of custom type variable
	fmt.Println("Students of section B : ")
	studentsB.getStudents()
}

// custom type declaration a variable
type classroom [] string

// receiver declaration of custom type variable.
func (d classroom) getStudents(){
	for i,student := range d {
		fmt.Println(i,student)
	}
}

