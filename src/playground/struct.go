package main

import (
	"log"
)

func main() {

	type Test struct {
		Ken    string
		Master int
	}

	x := Test{
		Ken:    "master",
		Master: 123,
	}
	// x := map[string]interface{}{
	// 	"ken":  "master",
	// 	"ken2": 123,
	// }

	for _, v := range x {
		log.Printf("%v", x.Ken)

	}

}
