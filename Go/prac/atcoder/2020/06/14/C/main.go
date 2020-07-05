package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var x, n int
	fmt.Scanf("%d %d", &x, &n)

	sc.Split(bufio.ScanWords)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	for i:=0; i < 100; i++ {
		if !contains(a, x-i) {
			fmt.Println(x-i)
			return
		}
		if !contains(a, x+i) {
			fmt.Println(x+i)
			return
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}


func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}