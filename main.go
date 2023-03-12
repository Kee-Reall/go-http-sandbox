package main

import (
	"fmt"
)

func main() {
	var st int = 2
	fmt.Println(st)
	addTwo(&st)
	fmt.Println(st)
	crt()
}

func addTwo(num *int) {
	fmt.Println(num)
	fmt.Println(&num)
	*num = *num + 2
	cursorPrint(&num)
}

func cursorPrint(curs **int) {
	fmt.Println(curs)
	fmt.Println(&curs)
}

func crt() {
	var i int16 = 0
	for true {
		print(i)
		if i == 32 {
			break
		}
	}
}
