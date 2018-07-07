package main

import (
	"time"
	"fmt"
)

func main() {
	// Unbuffered Channel: synchronous
	// Once a goroutine puts something in the box,
	// no other goroutine can put anything.
	c := make(chan int)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)

	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers2(w1)
	go printLetters2(w2)
	// They are just there to unblock the program once the goroutines complete

	<- w1 // trying to remove some value from the channel w1. Block until something is available
	<- w2 // trying to remove some value from the channel w2

	// Buffered Channel: FIFO
	// A goroutine can continually add things into this box w/o blocking
	// until there is no more space in the buffer.
	// another goroutine can  continually remove things from this box.
	c = make(chan int, 2)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}

func printNumbers2(w chan bool) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d \n", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A' + 100; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c \n", i)
	}
	w <- true
}

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <- c
		fmt.Println("Caught <<", num)
	}
}