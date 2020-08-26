package main

import "testing"

func TestNewStudents(t *testing.T){
	s := initStudents()
	if len(s) != 4 {
		t.Errorf("Expected length is 4 but the resultant length is %v", len(s))
	}
}
