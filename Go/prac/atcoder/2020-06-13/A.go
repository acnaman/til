package main

import (
	"fmt"
)

func mainA() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Printf("%s\n", s[0:3])
}
