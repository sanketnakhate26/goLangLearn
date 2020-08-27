package main

import "fmt"

type class struct {
	section string
	std int
}

type student struct {
	firstName string
	lastName string
	enroll class
}

func main(){
	s := student{"Rohit", "Sharma", class{"A",10}}
	fmt.Printf("%+v",s)
}
