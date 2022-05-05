package main

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	val int
}

//constructor
func NewCounter() *Counter {
	return &Counter{
		val: 0,
	}
}

func (c *Counter) inc() {
	c.Lock()
	c.val++
	c.Unlock()
}
 
func main() {

	var wg sync.WaitGroup

	counter := NewCounter()

	countGoroutines := 150

	wg.Add(countGoroutines)
	for i := 0; i < countGoroutines; i++ {
		go func () {
			defer wg.Done()
			counter.inc()
		} ()
	}

	wg.Wait()

	fmt.Println(counter.val) //print countGoroutines

}