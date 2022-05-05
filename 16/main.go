package main

/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/

import (
	"fmt"
	"math/rand"
)

// https://ru.wikipedia.org/wiki/Быстрая_сортировка
func qsort (arr []int, isIncrease bool) []int {
	if len(arr) < 2 { return arr }

	left, right := 0, len(arr) - 1

	pivotIndex := rand.Int() % len(arr)

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	if isIncrease {
		for i := range arr {
			if arr[i] < arr[right] {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}
	} else {
		for i := range arr {
			if arr[i] > arr[right] {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	qsort(arr[:left], isIncrease)
	qsort(arr[left + 1:], isIncrease)

	return arr
}

func main() {

	arr := []int{ -4, -9, 2, 4, 1, 0, 2, -5, 8, -3, 10, 9}

	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] < arr[j]
	// })

	arr = qsort(arr, false)
	fmt.Println(arr)

	arr = qsort(arr, true)
	fmt.Println(arr)

}