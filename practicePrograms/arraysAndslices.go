package main

import "fmt"

func main(){
	// Array declaration i.e. predefine list limit. Array is immutable list
	cards_array := [5] string {"Ace of club"}
	// Slice declaration i.e. without any list limit. Slice is mutable list
	cards_slice := [] string {"Ace of Heart"}

	// List assignment
	cards := append(cards_slice, "5 of Heart")

	// List iteration
	for i := range cards {
		fmt.Println(cards[i])
	}

	// Printing whole list
	fmt.Println("Cards_Array : ",cards_array)
	fmt.Println("Cards_Alice : ",cards_slice)
}