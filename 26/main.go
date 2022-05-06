package main

/*Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false*/

import (
	"fmt"
	"os"
	"strings"
)

func Uniq(str string) bool {
	str = strings.ToLower(str)
	runes := []rune(str)
	set := make(map[rune]bool)
	for _, val := range runes {
		set[val] = true
	}
	for _, val := range runes {
		if set[val] {
			set[val] = false 
		} else {
			return false
		}
	}
	return true
}

func main() {
	str := os.Args[1]

	fmt.Println(Uniq(str))
}