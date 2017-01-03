package main

import (
	"log"
)

func main() {
	a := []int{1, 2, 3}
	sample(a...)
}

func sample(a ...int) {

	for _, i := range a {
		log.Println(i)
	}

}
