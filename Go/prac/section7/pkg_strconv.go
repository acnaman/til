package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseintfunc()
}

func parseintfunc() {
	fmt.Println(strconv.ParseInt("123", 10, 0))
	fmt.Println(strconv.ParseInt("-1234", 10, 0))
	fmt.Println(strconv.ParseInt("740", 8, 0))
	fmt.Println(strconv.ParseInt("FFDD", 16, 0))

}

func strconvfunc() {
	b := true
	fmt.Printf("%T\n", strconv.FormatBool(b))
}

func parseboolfunc() {
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("0"))
}
