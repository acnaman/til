package main

import (
	"fmt"
	"math"
)

func main() {

	const (
		X = 34
		Y = "test"
	)

	print("Hello World")

	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

	fmt.Printf("value = %v\n", 1.0000000000000000)
	fmt.Printf("value = %v\n", 1.0000000000000001)
	fmt.Printf("value = %v\n", 1.0000000000000002)
	fmt.Printf("value = %v\n", 1.0000000000000003)
	fmt.Printf("value = %v\n", 1.0000000000000004)
	fmt.Printf("value = %v\n", 1.0000000000000005)
	fmt.Printf("value = %v\n", 1.0000000000000006)
	fmt.Printf("value = %v\n", 1.0000000000000007)
	fmt.Printf("value = %v\n", 1.0000000000000008)

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}
	fmt.Printf("a1代入前 %v\n", a1)
	a1 = a2
	fmt.Printf("a1代入後 %v\n", a1)

	q, r := div(43, 5)
	fmt.Printf("商：%d / 剰余：%d\n", q, r)

	f := returnFunc()
	f()

	f2 := later()
	fmt.Println(f2("Golang"))
	fmt.Println(f2("is"))
	fmt.Println(f2("awesome!"))
	fmt.Printf("X:%v, Y:%v\n", X, Y)

	if _, err := testFunc(); err != nil {
		fmt.Println("error occured!")
	} else {
		fmt.Println("testfunc doesn't return error!")
	}
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func returnFunc() func() {
	return func() {
		fmt.Println("retrunfunc test")
	}
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func testFunc() (string, error) {
	return "test", nil
}
