package main

import (
	"fmt"
	"time"
)

func main() {
}

func printNumbers1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A' + 10; i ++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	// Ends before goroutine can output to std
	go printNumbers1()
	go printLetters1()
}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}