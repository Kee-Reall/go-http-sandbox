package main

import (
	"fmt"
	"log"
	"sandbox/five"
)

func main() {
	for i := 1; i < 10; i++ {
		res, err := five.Five(4)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("jmmy")
		fmt.Print(res)
	}
}
