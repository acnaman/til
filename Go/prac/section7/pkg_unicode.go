package main

import (
	"fmt"
	"unicode"
)

func isXXfunc() {
	fmt.Println(unicode.IsDigit('X'))
	fmt.Println(unicode.IsDigit('3'))
	fmt.Println(unicode.IsDigit('３'))
	fmt.Println(unicode.IsLetter('3'))
	fmt.Println(unicode.IsLetter('あ'))
	fmt.Println(unicode.IsSpace(' '))
	fmt.Println(unicode.IsSpace('\t'))
	fmt.Println(unicode.IsSpace('_'))
	fmt.Println(unicode.IsControl('\r'))
	fmt.Println(unicode.IsControl('\n'))
	fmt.Println(unicode.IsMark('〒'))
}
