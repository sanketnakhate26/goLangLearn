package main

import "fmt"

type student struct {
	firstName string
	lastName string
}

func main(){
	rohit := student{"Rohit" , "Sharma"}
	var unknown student
	fmt.Print(rohit)
	fmt.Println(unknown)

}
