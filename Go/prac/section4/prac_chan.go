package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	func4()
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

func func4() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 10)

	ch1 <- 1
	ch2 <- 2

LOOP:
	for {
		select {
		case <-ch1:
			fmt.Println("------ch1から受信------")
		case <-ch2:
			fmt.Println("------ch2から受信------")
		case ch3 <- 3:
			fmt.Println("------ch3へ送信------")
		default:
			fmt.Println("受信送信できません")
			break LOOP
		}
	}
}

func func5() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			i := <-ch1
			ch2 <- i * 2
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- i - 1
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println("sub", i)
	}
}
