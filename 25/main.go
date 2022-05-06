package main

/*Реализовать собственную функцию sleep.*/

import (
	"fmt"
	"time"
)

//wait "duration" seconds and print
func Sleep_After(duration int) {
	fmt.Println(<- time.After(time.Duration(duration) * time.Second))
}

func Sleep_Tick(duration int) {
	fmt.Println(<- time.Tick(time.Duration(duration) * time.Second))
}

func main() {
	fmt.Println(time.Now())
	Sleep_After(2)
	Sleep_Tick(2)

	ticker := time.NewTicker(2 * time.Second)
	fmt.Println(<- ticker.C)
}