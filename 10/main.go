package main

/*Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.*/

import (
	"fmt"
	"math"
)

func main() {

	temperature := []float64{ -25.4, -27.0, 13.0, -10.1, -10.0, -0.1, 0, 0, 0.1, 10.0, 19.0, 15.5, 24.5, -21.0, 32.5 }
	step := 10

	subsets := make(map[int][]float64)

	//map: ..., -20:[-20; -10), -10:[-10; 0), 0:[0; 10), 10:[10; 20), ...
	for _, val := range temperature {
		div := val / float64(step)
		mod := math.Mod(val, float64(step))

		//check negative number and number / 10 != 0
		if div < 0 && mod != 0 { div -= 1 } //left border "["
		firstDigit := int(div)
		subsets[firstDigit*step] = append(subsets[firstDigit*step], val)
	}

	fmt.Println(subsets)
}