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
	c := make([]int, 26)

	for i := 0; i < 26; i++ {
		c[i] = nextInt()
	}

	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			s[i][j] = nextInt()
		}
	}

	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = nextInt() - 1 // X日目が配列とずれるから1引く
	}

	for i := 1; i <= n; i++ { // X日目が配列とずれるから1足す
		fmt.Println(culcSatisfuctionScore(c, s, t, i))
	}

}

func culcSatisfuctionScore(c []int, s [][]int, t []int, day int) int {
	var score int

	lastdi := make([]int, 26)
	for i := 0; i < day; i++ {
		score += s[i][t[i]]
		lastdi[t[i]] = i + 1
		for j := 0; j < 26; j++ {
			score -= c[j] * (i + 1 - lastdi[j])
		}
	}
	return score
}

func query 

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxIntSlice(a int[]) int {
	b := a[0]
	for _, v range(a)
	if b < v {
		b = v
	}
	return b
}


func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
