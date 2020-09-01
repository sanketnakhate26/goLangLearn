package main

import "fmt"

func main() {
	// map declaration
	marks := map[string]int64{
		"maths": 95,
		"phy":   96,
		"chem":  87,
	}

	printMap(marks)
}

// funcion to print map values
func printMap(m map[string]int64) {
	for key, value := range(m){
		fmt.Println(value," marks got in ",key)
	}
}
