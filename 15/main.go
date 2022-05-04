package main

/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

import (
	"fmt"
	"math/rand"
)

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
  fmt.Println(justString)

  v1 := createHugeString_2(1 << 11)
  justString = string(v1[:100])
  fmt.Println(justString)
}

func createHugeString(size int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, size)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
}

func createHugeString_2(size int) []rune {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, size)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return str
}

func main() {
	someFunc()
}