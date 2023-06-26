package main

import (
	"fmt"
)

func main() {

	// create channels
	number := make(chan int)
	message := make(chan string)

	// function call with goroutine
	go channelNumber(number)
	go channelMessage(message)

	// selects and executes a channel
	select {
	case firstChannel := <-number:
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data:", secondChannel)

	// default case
	default:
		fmt.Println("Wait!! Channels are not ready for execution")
	}

}

// goroutine to send data to the channel
func channelNumber(number chan int) {

	// sleeps the process for 2 seconds
	//time.Sleep(2 * time.Second)

	number <- 15
}

// goroutine to send data to the channel
func channelMessage(message chan string) {

	// sleeps the process by 2 seconds
	//time.Sleep(2 * time.Second)

	message <- "Learning Go Select"
}
