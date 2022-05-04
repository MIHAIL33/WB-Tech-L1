package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/


func main() {
	//read number from args
	number, err := strconv.Atoi((os.Args[1])) 
	if err != nil {
		panic(err)
	}

	//read i-th bit from args
	i, err := strconv.Atoi((os.Args[2])) 
	if err != nil {
		panic(err)
	}

	if i > 64 || i < 1 {
		panic("i must be between 1 and 64")
	}

	//get all bits with zeroes
	val := fmt.Sprintf("%064b", number)
	fmt.Println(val)
	bits := strings.Split(val, "")

	//replace
	if bits[len(bits) - i] == "0" {
		bits[len(bits) - i] = "1"
	} else {
		bits[len(bits) - i] = "0"
	}

	val = strings.Join(bits, "")
	fmt.Println(val)
	
	//binary string to int
	res, _ := strconv.ParseInt(val, 2, 64)
	fmt.Println(res)
}