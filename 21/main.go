package main

/*Реализовать паттерн «адаптер» на любом примере.*/

import "fmt"

type assembly interface { 
	insertEngine()
}

type Auto struct {}

func (a *Auto) insertV12(as assembly) {
	as.insertEngine()
}

type EngineV12 struct {}

func (e *EngineV12) insertEngine() {
	fmt.Println("EngineV12")
}

type EngineV8 struct {}

func (e *EngineV8) addEngine() {
	fmt.Println("EngineV8")
}

type EngineV8Adepter struct {
	enginev8 *EngineV8
}

func (e *EngineV8Adepter) insertEngine() {
	e.enginev8.addEngine()
}


func main() {

	auto := &Auto{}
	engineV12 := &EngineV12{}
	auto.insertV12(engineV12)

	engineV8 := &EngineV8{}
	engineV8Adepter := &EngineV8Adepter{
		enginev8: engineV8,
	}

	auto.insertV12(engineV8Adepter)

}