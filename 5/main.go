package main

/*Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.*/

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	seconds, err := strconv.Atoi((os.Args[1])) //get seconds from args
	if err != nil {
		panic(err)
	}

	exitChannel := make(chan bool, 1) //channel for exit
	defer close(exitChannel)

	//run function after timer expires
	timer := time.AfterFunc(time.Second * time.Duration(seconds), func() {
		exitChannel <- true
	})
	defer timer.Stop() //stop timer

	channel := make(chan string, 1) //channel for data
	defer close(channel)

	for {
		select {
		case channel <- "data": //writing "data" to channel
		case data := <- channel: //reading "data" from channel
			fmt.Println(data) //print "data"
		case <- exitChannel: // get signal of exit
			fmt.Println("Finish")
			return
		}
	}

}