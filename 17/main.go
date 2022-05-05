package main

/*Реализовать бинарный поиск встроенными методами языка.*/

import "fmt"

// https://ru.wikipedia.org/wiki/Двоичный_поиск
func binarySearch(val int, arr []int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2
		
		if arr[median] < val {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != val {
		return false
	}

	return true
}

func main() {

	arr := []int{ -3, -2, 1, 5, 6, 8, 9, 10 }

	fmt.Println(binarySearch(-3, arr)) //true
	fmt.Println(binarySearch(8, arr)) //true
	fmt.Println(binarySearch(0, arr)) //false

}