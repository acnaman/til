package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc.Split(bufio.ScanWords)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	count := 0

	for i := 0; i < n; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if a[i]%a[j] == 0 {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}

	fmt.Println(count)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
