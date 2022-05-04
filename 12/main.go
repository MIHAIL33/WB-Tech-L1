package main

/*Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.*/

import "fmt"

func unique(slice []string) (uniq []string) {
	hash := make(map[string]bool)
	for _, val := range slice {
		hash[val] = true //plug
	}
	for key := range hash {
		uniq = append(uniq, key)
	}
	return //return uniq
}

func main() {

	slice := []string{ "cat", "cat", "dog", "cat", "tree", "lake", "tree"}

	result := unique(slice)

	fmt.Println(result)

}