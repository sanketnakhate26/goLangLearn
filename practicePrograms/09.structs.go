package main

import "fmt"
g
type student struct {
	firstName string
	lastName string
}

func main(){
	rohit := student{"Rohit" , "Sharma"}
	var unknown student
	rohit.print()
	unknown.print()

}

func (s student) print() {
	fmt.Printf("%+v",s)
}

