package main

/*Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.*/

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	arr := [5]int{ 2, 4, 6, 8, 10 }
	sum := 0


	//Version 1 with channel

	// sumChannel := make(chan int, 5)

	// for _, val := range arr {
	// 	wg.Add(1)
	// 	go func(val int) {
	// 		defer wg.Done()
	// 		sumChannel <- int(math.Pow(float64(val), 2))
	// 	} (val)

	// 	// addition without blocking for
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		sum += <- sumChannel
	// 	} ()

	// 	// sum += <- sumChannel  - simple version with blocking for
	// }

	//Version 2 with mutex

	var mu sync.Mutex

	for _, val := range arr {
		wg.Add(1)
		go func (sum *int, val int)  {
			defer wg.Done()
			mu.Lock() //Lock sum
			*sum += int(math.Pow(float64(val), 2))
			mu.Unlock() //Unlock sum
		} (&sum, val)
	}

	wg.Wait() // wait all goroutines

	fmt.Println(sum) //print 220
}