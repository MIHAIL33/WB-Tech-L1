package main 

import (
	"fmt"
)

type Human struct {
	age int
}

func (h Human) happyBirthday() int {
	h.age++
	return h.age 
}

//Version 1
type Action struct {
	person Human
}

func (a Action) happyBirthday() int {
	return a.person.happyBirthday()
}

//Version 2
type Action2 struct {
	Human
}

func main() {
	action := Action{person: Human{age: 10}}
	fmt.Println(action.happyBirthday()) //Print 11

	action2 := Action2{Human{age: 15}}
	fmt.Println(action2.happyBirthday()) //Print 16
}