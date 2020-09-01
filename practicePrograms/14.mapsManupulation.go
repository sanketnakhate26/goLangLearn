package main

import "fmt"

func main() {
	// map declaration
	marks := map[string]int64{
		"maths": 95,
		"phy":   96,
		"chem":  87,
	}

	fmt.Println("marks before update :")
	fmt.Println(marks)

	fmt.Println("marks after update (new value added) :")
	marks["computer sci"]=92
	fmt.Println(marks)

	fmt.Println("marks after update (new value deleted)")
	delete(marks,"computer sci")
	fmt.Println(marks)


}