package main

import "fmt"

func main() {
	mfunc1()
}

func mfunc1() {
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"
	fmt.Println(m)

	m2 := map[int]string{
		1:  "US",
		81: "Japan",
		86: "China",
	}
	fmt.Println(m2)

	data, ok := m2[2]
	fmt.Println(data, ok)

	for k, v := range m {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
}
