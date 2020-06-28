package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func mainA() {
	sc.Split(bufio.ScanWords)

	for i := 0; i < 5; i++ {
		x := nextInt()
		if x == 0 {
			fmt.Println(i + 1)
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
