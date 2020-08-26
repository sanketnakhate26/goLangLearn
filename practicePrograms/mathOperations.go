package main

import (
	"fmt"
	"math/rand"
	"time"
)

type numbers []string
func main(){
	n :=initNumbers()
	fmt.Println("Numbers initial sequence :")
	n.print()
	n.shuffle()
	fmt.Println("Numbers after shuffle :")
	n.print()
}

func (n numbers) print(){
	fmt.Println(n)
}

func initNumbers() numbers{
	return []string{"0","1","2","3","4","5"}
}

func (n numbers) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i,_ := range n{
		newposition := r.Intn(len(n)-1)
		n[i], n[newposition] = n[newposition] , n[i]
	}
}
