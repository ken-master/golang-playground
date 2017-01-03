package main

import (
	"fmt"
)

func main() {

	defer willCatch()

	willPanic()

}

func willPanic() {

	fmt.Println("ken")
	fmt.Println("master")
	panic("THIS IS SPAAAARRTAA!!")

}

func willCatch() {

	if recover() != nil {
		fmt.Println("DONT PANIC IT'S ORGANIC")
	}
}
