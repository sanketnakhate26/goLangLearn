// declaring main package
package main

// importing fmt package
import "fmt"

// declaring main function
func main(){
	noOfCards := countCards()
	card := selectedCard()
	piValue := getPiValue()

	fmt.Println("No of Cards = ",noOfCards," and the selected card is ",card)
	fmt.Println("Value of Pi = ",piValue)
}
// func <function name> <return type> is the syntax for declaring a function. This is a function to return a int variable
func countCards() int {
	return 52
}
// function to return a string variable
func selectedCard() string {
	return "King of Heart"
}
// function to return a float value
func getPiValue() float64{
	return 3.14
}
