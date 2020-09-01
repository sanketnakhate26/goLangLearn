package main

import "fmt"

type student struct {
	firstName string
	lastName string
}

func main(){
	student1 := student{"Rohit" , "Sharma"}
	pointerToStudent1 := &student1
	fmt.Println("Value of object is :  ",*pointerToStudent1)
	pointerToStudent1.update("Sanket")
	// Value of object is changed. By defalut the parameters are pass as pass by value in golang. Pointers can be used to modify the value.
	fmt.Println("Value of object after update is : ",*pointerToStudent1)

}

func (s student) print() {
	fmt.Printf("%+v",s)
}

func (newPointer *student) update(newname string) *student {
	(*newPointer).firstName = newname
	return newPointer
}

