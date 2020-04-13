package main

import "fmt"

func main() {
	arrayToSlice()
}

func afunc() {
	s := make([]int, 10)
	fmt.Println(s)

	s = []int{10, 3, 3, 3, 3, 3, 33, 3}
	fmt.Println(s)

	s = append(s, 5, 6, 6, 7)
	fmt.Println(s)
}
func bfunc() {
	s := make([]int, 0, 0)
	fmt.Printf("(1) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("(2) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, []int{2, 3, 4}...)
	fmt.Printf("(3) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 5)
	fmt.Printf("(4) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 6, 7, 8, 9)
	fmt.Printf("(5) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 5)
	fmt.Printf("(6) len=%d, cap=%d\n", len(s), cap(s))

	s = make([]int, 1024, 1024)
	fmt.Printf("(8) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 0)
	fmt.Printf("(9) len=%d, cap=%d\n", len(s), cap(s))
}

func sum(s ...int) int {
	n := 0

	for _, v := range s {
		n += v
	}
	return n
}

func arrayToSlice() {
	a := [3]int{1, 2, 3}
	s := a[:]
	fmt.Printf("(1) len=%d, cap=%d\n", len(s), cap(s))

	s = append(s, 4)
	fmt.Printf("(2) len=%d, cap=%d\n", len(s), cap(s))

	a[0] = 9
	fmt.Printf("(3) slice=%v\n", s)
	fmt.Printf("(4) array=%v\n", a)

}
