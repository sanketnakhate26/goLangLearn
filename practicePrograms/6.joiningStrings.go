package main

import (
	"fmt"
	//"io/ioutil"
	"strings"
)

type classstudents []string

func main(){
	classroom := getclass()
	// Join function to join the list of string with the given seperator
	class := strings.Join(classroom,",")
	fmt.Println(class)
}

func getclass()classstudents{
	c := classstudents{}
	c = append(c,"one")
	c = append(c,"two")

	return c
}




