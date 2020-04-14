package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	func3()
}

func func1() {
	ch := make(chan int)

	go receiver(ch)

	for i := 0; i < 20; i++ {
		ch <- i
		fmt.Println("main", i)
	}
}

func func2() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	i, ok := <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done")
}

func func3() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	fmt.Println("numgoroutine:", runtime.NumGoroutine())

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println("sub", i)
	}
}
