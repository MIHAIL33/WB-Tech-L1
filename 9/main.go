package main

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны
выводиться в stdout.*/

import (
	"fmt"
	"sync"
)

//writing x to firstChannel
func Write(x []int, firstChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, val := range x {
		firstChannel <- val
	}
	close(firstChannel)
}

//reading x from firstChannel and writing x^2 to secondChannel
func SqWrite(firstChannel chan int, secondChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range firstChannel {
		secondChannel <- val * 2 
	}
	close(secondChannel)
}

//reading x^2 from secondChannel and writing to stdout 
func Read(secondChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range secondChannel {
		fmt.Println(val)
	}
}

func main() {

	var wg sync.WaitGroup 

	x := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }

	firstChannel := make(chan int) // for x
	secondChannel := make(chan int) // for x^2

	wg.Add(3)
	go Write(x, firstChannel, &wg)
	go SqWrite(firstChannel, secondChannel, &wg)
	go Read(secondChannel, &wg)

	wg.Wait() // wait all goroutines
}