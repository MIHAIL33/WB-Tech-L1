package main

/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	str := "жваотывж3х10хоаваф,23123=犬ш=ш495г4гзтеолр459j;ldfpahpahgp"

	fmt.Println(reverse(str))

}