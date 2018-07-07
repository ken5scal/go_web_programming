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

	// Select allows you to select one of many channels to receive from or send to.
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond) // adding this because otherwise `select` is called too quickly
		select {
		case msg:= <- a:
			fmt.Printf("%s from A\n",msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			// if you do not set this, deadlock occurs because
			// both a and b has already sent value and received, thus blocked and asleep.
			fmt.Println("Default")
		}
	}

	a, b = make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 { //ok1 and ok2 become false when channels close
		select {
		case msg, ok1 = <- a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d \n", i)
	}
	w <- true
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A' + 10; i++ {
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

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Good night World!"
	close(c)
}
