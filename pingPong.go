package main

import (
	"fmt"
	"time"
)

func createMessage() <-chan string {
	message := make(chan string)
	go func() {
		for i := 0; ; i++ {
			if i%2 == 0 {
				message <- "Ping"
			} else {
				message <- "Pong"
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return message
}

func printMessage(message <-chan string) {
	for msg := range message {
		fmt.Println(msg)
	}
}

func main() {
	printMessage(createMessage())
}
