package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	t := nextInt()

	for i := 0; i < t; i++ {
		testcase()
	}

}

func testcase() {
	n := nextInt()
	a := make([]int, n)
	tmp := strings.Split(nextString(), " ")

	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(tmp[i])
	}
	s := nextString()

	target := 0

	for i := n - 1; i > 0; i-- {
		si := s[i]
		if si == '0' {
			target = a[i] ^ target
		} else {
			target = a[i] &^ target
		}
	}

	if target == a[0] {
		fmt.Println("0")
	} else {
		fmt.Println("1")
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}