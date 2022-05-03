package main

/*Реализовать все возможные способы остановки выполнения горутины.*/

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	//Channel Version
	exitChannel := make(chan bool)
	defer close(exitChannel)

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for {
			select {
			case <- exitChannel:
				return
			default:
				fmt.Println("I'm first goroutine")
			}
		}
	} ()

	time.Sleep(2 * time.Second)
	exitChannel <- true 

	//Context Version
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func (ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				return
			default:
				fmt.Println("I'm second goroutine")
			}
		}
	} (ctx)

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
}