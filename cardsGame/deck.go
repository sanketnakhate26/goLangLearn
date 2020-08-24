package main

import "strings"

type deck []string

func newDeck () deck {

	cards := deck{}
	cardSymbols := []string{"Club", "Spade", "Heart","Diamond"}
	cardValues := []string{"Ace","One","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}


	for _,symbol := range cardSymbols{
		for _,value := range cardValues{
			cards = append(cards,value+" of "+symbol)
		}
	}

	return cards
}
func deal (d deck, handsize int) (deck, deck){
	return d[:handsize], d[handsize:]
}

func show (d deck) [] string {
	cardList := [] string{}
	for _,card := range d {
		cardList = append(cardList,card)
	}
	return cardList
}
func (d deck) toString() string{
	return strings.Join([]string(d),",")
}