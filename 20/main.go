package main

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

import (
	"fmt"
	"strings"
)

func reverse(str []string) []string {
	for i, j := 0, len(str) - 1; i < j; i, j = i + 1, j - 1 {
		str[i], str[j] = str[j], str[i]
	}
	return str
}

func main() {

	str := "dog car snow sun tree дом"

	strs := strings.Split(str, " ")

	strs = reverse(strs)

	str = strings.Join(strs, " ")

	fmt.Println(str)

}