package main

/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала
и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.*/

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	//get count (int) workers from args
	countWorkers, err := strconv.Atoi((os.Args[1])) 
	if err != nil {
		panic(err)
	}

	workChannel := make(chan string, countWorkers) //channel for write and read data
	defer close(workChannel)
	exitSignal := make(chan os.Signal, 1) //channel for catching signal
	defer close(exitSignal)
	exitWorker := make(chan bool, countWorkers) //channel
	defer close(exitWorker)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup //wait all goroutines

	// start goroutines, which reading
	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go func () {
			defer wg.Done()
			for {
				select {
				case data := <-workChannel:
					fmt.Fprintln(os.Stdout, data) // Print "data"
				case <- exitWorker:
					fmt.Println("Worker finished")
					return
				}
			}
		} ()
	}

	for {	
		select {
		case workChannel <- "data": //writing "data" in channel 
		case <-exitSignal: //get signal (ctrl + c)
			for i := 0; i < countWorkers; i++ { 
				exitWorker <- true //post all goroutines to end
			}
			wg.Wait() //wait all goroutines
			fmt.Println("Finish")
			return 
		}
	}
}