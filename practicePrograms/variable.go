package main

import "fmt"

func main(){
	// declaration of a variable. <var> <var name> <var type> = <value>. Go is statically typed language
	var card1 string = "Ace of spade"
	// other way of declaring variable. Its kinda short hand variable declaration
	card2 := "Ace of diamond"
	card3 := "None"
	// variable value assign
	card3 = card2

	fmt.Println("card1 = ",card1," and card2 = ",card2," and card3 = ",card3)
}
