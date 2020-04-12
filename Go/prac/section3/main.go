package main

import (
	"fmt"
	"math"
)

func main () {
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

	a1 := [3]int{1,2,3}
	a2 := [3]int{4,5,6}
	fmt.Printf("a1代入前 %v\n", a1)
	a1 = a2
	fmt.Printf("a1代入後 %v\n", a1)

	q, r := div(43, 5)
	fmt.Printf("商：%d / 剰余：%d\n", q, r)
}

func div (a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
