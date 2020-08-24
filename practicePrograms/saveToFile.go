package main

import "fmt"
import "io/ioutil"
import "strings"

type students []string
func main(){
	stu := initStudents()
	fmt.Println(stu)
	stu.saveToFile("new.txt")

}
func initStudents() students{
	return []string{"Sanket","Raj","Rohit","Gargi"}
}

func (s students) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(s.toString()),0666)
}

func (s students) toString() string{
	return strings.Join([]string(s),",")
}