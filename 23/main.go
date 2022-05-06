package main

/*Удалить i-ый элемент из слайса.*/

import "fmt"

//Version 1
func RemoveElementByIndex(slice []int, i int) []int {
	slice[i] = slice[len(slice) - 1] //last element instead of removed
	slice = slice[:len(slice) - 1] //trim slice
	return slice
} 

//Version 2
func RemoveElementByIndex2(slice []int, i int) []int {
	copy(slice[i:], slice[i + 1:]) //shift left by one index
	slice = slice[:len(slice) - 1] //trim slice
	return slice
} 

func main() {

	slice := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	fmt.Println(slice)

	slice = RemoveElementByIndex(slice, 2)
	fmt.Println(slice)

	slice = RemoveElementByIndex2(slice, 4)
	fmt.Println(slice)

}