package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	funcb()
}

func funca() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d\n", a*b)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func funcb() {
	var n int
	const MAX_AMOUNT = 1e+18
	fmt.Scanf("%d", &n)

	sc.Split(bufio.ScanWords)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		if a[i] == 0 {
			fmt.Println("0")
			return
		}
	}

	var v int64 = 1
	for i := 0; i < n; i++ {
		if a[i] <= MAX_AMOUNT/v {
			v *= a[i]
		} else {
			fmt.Println("-1")
			return
		}
	}

	fmt.Printf("%d\n", v)
}

func funcc() {
	var a float64
	var b float64
	fmt.Scanf("%f %f", &a, &b)
	fmt.Printf("%d\n", int(a*b))
}

func funcd() {
	var n int
	fmt.Scanf("%d", &n)

	count := 0
	for i := 1; i <= n; i++ {
		//isZ
		if n%i == 0 {
			count++
			n /= i
		}
	}
	fmt.Printf("%d\n", count)
}

/*func isZ (num int) bool {
  for i:=0; i<=num;
}*/

func funce() {
	var a int
	fmt.Scanf("%d", &a)
	fmt.Printf("%d\n", a)
}
