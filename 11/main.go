package main

/*Реализовать пересечение двух неупорядоченных множеств.*/

import "fmt"

func insersection(firstSlice, secondSlice []int) (inter []int) {
	hash := make(map[int]bool)
	for _, val := range firstSlice {
		hash[val] = true
	}
	for _, val := range secondSlice {
		if hash[val] {
			inter = append(inter, val)
			hash[val] = false //lock same number
		}
	}
	return //return inter
}

func main() {

	firstSlice := []int{ -1, 2, -5, 0, 8, 10, 15, -16 }
	secondSlice := []int{ -2, -4, 0, 0, 22, 10, -16, 20 }

	result := insersection(firstSlice, secondSlice)

	fmt.Println(result)

}