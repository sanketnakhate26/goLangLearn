package main

import "fmt"

type student struct {
	firstName string
	lastName string
}

func main(){
	student1 := student{"Rohit" , "Sharma"}
	fmt.Println("Value of object is :  ",student1)
	student1.update("Sanket")
	// Value of objectt did not changed. By defalut the parameters are pass as pass by value in golang
	fmt.Println("Value of object after update is : ",student1)

}

func (s student) print() {
	fmt.Printf("%+v",s)
}

func (s student) update(newname string) student {
	s.firstName = newname
	return s
}
