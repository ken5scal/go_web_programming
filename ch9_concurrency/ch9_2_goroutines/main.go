package main

import (
	//"fmt"
	"time"
	"sync"
)

func main() {
}

func printNumbers1() {
	for i := 0; i < 10; i++ {
	//	fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A' + 10; i ++ {
	//	fmt.Printf("%c ", i)
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

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A' + 100; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%c ", i)
	}
	wg.Done()               // Decrement Counter(If you forget, dead lock occurs)
}

func print2() {
	var wg sync.WaitGroup
	wg.Add(2)
	printNumbers2(&wg)
	printLetters2(&wg)
	wg.Wait()               // Decrement Counter(If you forget, dead lock occurs)
}

func goPrint2() {
	var wg sync.WaitGroup   // Declare Wait group (Ensure all goroutines complete before moving on to the next thing)9
	wg.Add(2)         // Set up counter
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()               // Blocks until counter reaches 0
}

//% go test -run x -bench . -cpu=1                                                       (git)-[master]
//goos: darwin
//goarch: amd64
//pkg: github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines
//BenchmarkPrint1         100000000               11.6 ns/op
//BenchmarkGoPrint1        1000000              1058 ns/op
//BenchmarkPrint2             2000            860882 ns/op
//BenchmarkGoPrint2        1000000             14028 ns/op
//PASS
//ok      github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines 18.759s

//% go test -run x -bench . -cpu=2                                                       (git)-[master]
//goos: darwin
//goarch: amd64
//pkg: github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines
//BenchmarkPrint1-2       100000000               11.5 ns/op
//BenchmarkGoPrint1-2      5000000               397 ns/op
//BenchmarkPrint2-2           2000            837545 ns/op
//BenchmarkGoPrint2-2      1000000             13644 ns/op
//PASS
//ok      github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines 19.333s

//% go test -run x -bench . -cpu=3                                                       (git)-[master]
//goos: darwin
//goarch: amd64
//pkg: github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines
//BenchmarkPrint1-3       100000000               11.8 ns/op
//BenchmarkGoPrint1-3      5000000               304 ns/op
//BenchmarkPrint2-3           2000            866585 ns/op
//BenchmarkGoPrint2-3       300000             12320 ns/op
//PASS
//ok      github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines 8.717s

//% go test -run x -bench .
//goos: darwin
//goarch: amd64
//pkg: github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines
//BenchmarkPrint1-4       100000000               11.4 ns/op
//BenchmarkGoPrint1-4      5000000               307 ns/op
//BenchmarkPrint2-4           2000            888034 ns/op
//BenchmarkGoPrint2-4       200000              9971 ns/op
//PASS
//ok      github.com/ken5scal/go_web_programming/ch9_concurrency/ch9_2_goroutines 19.198s