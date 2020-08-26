package main

import (
	"fmt"
	"os"
)
import "io/ioutil"
import "strings"

type students []string
func main(){
	stu := initStudents()
	fmt.Println("File contents before save : ")
	stu.print()
	fmt.Println("Saving file.....")
	stu.saveToFile("new.txt")
	fmt.Println("Reading contents from the saved file :")
	newStudents :=newStudentsFromFile("new.txt")
	newStudents.print()

}
func initStudents() students{
	return []string{"Sanket","Raj","Rohit","Gargi"}
}

func (s students) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(s.toJoinString()),0666)
}

func (s students) toJoinString() string{
	return strings.Join([]string(s),",")
}

func newStudentsFromFile(filename string) students {
	bs, err := ioutil.ReadFile("new.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
		}
	s := strings.Split(string(bs),",")
	return students(s)
}

func (s students) print() {
	fmt.Println(s)
}