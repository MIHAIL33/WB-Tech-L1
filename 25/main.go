package main

/*Реализовать собственную функцию sleep.*/

import (
	"fmt"
	"time"
)

//wait "duration" seconds and print
func Sleep(duration int) {
	fmt.Println(<- time.After(time.Duration(duration) * time.Second))
} 

func main() {
	fmt.Println(time.Now())
	Sleep(5)
}