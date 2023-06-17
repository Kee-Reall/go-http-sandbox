package main

import (
	"fmt"
	"os/exec"
)

func main() {
	defer handlePanic()
	var cmd string = "node"
	is := checkCmd(cmd)
	var ableStr string
	if is {
		ableStr = "available"
	} else {
		ableStr = "not available"
	}
	fmt.Printf("%s is %s", cmd, ableStr)
}

func checkCmd(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%v+", err.Error())
	}
	return err == nil
}

func handlePanic() {
	if msg := recover(); msg != nil {
		fmt.Printf("%v+", msg)
	}
}
