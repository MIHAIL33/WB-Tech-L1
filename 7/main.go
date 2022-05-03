package main

/*Реализовать конкурентную запись данных в map.*/

import (
	"fmt"
	"strconv"
	"sync"
)

// Version 1 with read lock
type MapWithMutex struct {
	sync.Mutex
	m map[string]int
}

//constructor
func NewMapWithMutex() *MapWithMutex { 
	return &MapWithMutex{
		m: make(map[string]int),
	}
}

//add new element
func (mw *MapWithMutex) Store(key string, value int) {
	mw.Lock()
	mw.m[key] = value
	mw.Unlock()
}

//get element by key
func (mw *MapWithMutex) Load(key string) (int, bool) {
	mw.Lock()
	defer mw.Unlock()
	val, ok := mw.m[key]
	return val, ok
}
 
// Version 2 without read lock
type MapWithMutex2 struct {
	sync.RWMutex
	m map[string]int
}

//constructor
func NewMapWithMutex2() *MapWithMutex2 {
	return &MapWithMutex2{
		m: make(map[string]int),
	}
}

//add new element
func (mw *MapWithMutex2) Store(key string, value int) {
	mw.Lock()
	mw.m[key] = value
	mw.Unlock()
}

//get element by key
func (mw *MapWithMutex2) Load(key string) (int, bool) {
	mw.RLock()
	defer mw.RUnlock()
	val, ok := mw.m[key]
	return val, ok
}

func main() {

	var wg sync.WaitGroup //wait all goroutines

	countGoruotines := 10 // count goroutines
	
	// Version 1 with Mutex
	myMap := NewMapWithMutex()

	//Writing
	for i := 0; i < countGoruotines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap.Store(strconv.Itoa(i), i)
		} (i)
	}

	wg.Wait()

	//Reading
	for i := 0; i < countGoruotines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, _ := myMap.Load(strconv.Itoa(i))
			fmt.Println(val)
		} (i)
	}

	fmt.Println(myMap)
	wg.Wait()

	// Version 2 with RWMutex
	myMap2 := NewMapWithMutex2()

	//Writing
	for i := 0; i < countGoruotines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap2.Store(strconv.Itoa(i), i)
		} (i)
	}

	wg.Wait()

	//Reading
	for i := 0; i < countGoruotines; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			val, _ := myMap2.Load(strconv.Itoa(i))
			fmt.Println(val)
		} (i)
	}

	fmt.Println(myMap2)
	wg.Wait()

	//Version 3 with sync.Map
	var myMap3 sync.Map

	//Writing
	for i := 0; i < countGoruotines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap3.Store(strconv.Itoa(i), i)
		} (i)
	}

	wg.Wait()

	//Reading
	for i := 0; i < countGoruotines; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			val, _ := myMap3.Load(strconv.Itoa(i))
			fmt.Println(val)
		} (i)
	}

	fmt.Println(&myMap3)
	wg.Wait()

}