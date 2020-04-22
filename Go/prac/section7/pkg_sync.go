package main

import (
	"fmt"
	"sync"
	"time"
)

func waitgroupfunc() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func(){
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done()
	}()

	wg.Wait()
}

var st struct{ A, B, C int }
var mutex *sync.Mutex

func UpdateAndPrintMutex(n int) {
	mutex.Lock()

	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st.A, st.B, st.C)

	mutex.Unlock()
}

func mutexfunc() {
	mutex = new(sync.Mutex)
	for {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
}

/* NG Pattern */
func UpdateAndPrint(n int) {
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st.A, st.B, st.C)
}

func callfunc() {
	for {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
}
