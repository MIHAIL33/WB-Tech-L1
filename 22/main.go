package main

/*Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.*/

import (
	"fmt"
	"math/big"
	"os"
)

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {

	var aStr, bStr, operator string

	fmt.Println("Enter first value:")
	_, err := fmt.Fscan(os.Stdin, &aStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter second value:")
	_, err = fmt.Fscan(os.Stdin, &bStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter operator (+, -, *, /)")
	_, err = fmt.Fscan(os.Stdin, &operator)
	if err != nil {
		panic(err)
	}

	//string to bigInt
	a := new(big.Int)
	a.SetString(aStr, 10)
	b := new(big.Int)
	b.SetString(bStr, 10)

	switch operator{
	case "+":
		fmt.Println(Add(a, b))
	case "-":
		fmt.Println(Sub(a, b))
	case "*":
		fmt.Println(Mul(a, b))
	case "/":
		fmt.Println(Div(a, b))
	}

}