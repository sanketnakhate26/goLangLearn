package main

import "fmt"

func main(){

	cardsList := newDeck()

	fmt.Println(cardsList)

	handCards, remainCards := deal(cardsList,5)

	fmt.Println("Cards in hand are : ",show(handCards))
	fmt.Println("Cards remain are : ",show(remainCards))
}