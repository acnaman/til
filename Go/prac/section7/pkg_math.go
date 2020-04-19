package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randsourcefunc() {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	fmt.Println(rnd.Intn(100))
	fmt.Println(rnd.Int())
}

func mathfunc() {

	fmt.Println(math.Abs(-14))
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Cbrt(8))
	fmt.Println("---Trunc---")
	fmt.Println(math.Trunc(1.5))
	fmt.Println(math.Trunc(-1.5))
	fmt.Println("---Floor---")
	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Floor(-1.5))
	fmt.Println("---Ceil---")
	fmt.Println(math.Ceil(1.5))
	fmt.Println(math.Ceil(-1.5))

}

func randfunc() {
	rand.Seed(256)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
}
