package main

/*Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.*/

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(2) //set count procs

	var wg sync.WaitGroup 

	arr := [5]int{ 2, 4, 6, 8, 10 }

	for _, val := range arr {
		wg.Add(1)
		go func (val int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout ,math.Pow(float64(val), 2)) //print in random order
		} (val)
	}

	wg.Wait() //wait all goroutines
}