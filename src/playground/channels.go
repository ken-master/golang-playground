package main

import (
	"log"
)

func main() {

	dataChannel := make(chan interface{})

	go func() {
		dataChannel <- 123
	}()

	log.Printf("%v", <-dataChannel)
}
