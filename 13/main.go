package main

/*Поменять местами два числа без создания временной переменной.*/

import "fmt"

func main() {
	
	a := 10 //1010
	b := 15 //1111

	a = a ^ b //0101 or 5
	b = b ^ a //1010 or 10
	a = a ^ b //1111 or 10

	fmt.Println(a)
	fmt.Println(b)
}