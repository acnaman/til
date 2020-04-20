package main

import (
	"fmt"
	"strings"
)

func main() {
	replacefunc()
}

func replacefunc() {
	fmt.Println(strings.Replace("AAAAA", "A", "X", 1))
	fmt.Println(strings.Replace("AAAAA", "A", "X", 2))
	fmt.Println(strings.Replace("AAAAA", "A", "X", -1))
	fmt.Println(strings.Replace("ABCDE", "A", "X", 2))
	fmt.Println(strings.Replace("C言語", "C", "Go", 1))
}

func repeatfunc() {
	fmt.Println("---- Repeat ----")
	fmt.Println(strings.Repeat("ABC", 3))
	fmt.Println(strings.Repeat("ABC", 0))
}

func containsfunc() {
	fmt.Println("---- Contains ----")
	fmt.Println(strings.Contains("ABCDE", "AB"))
	fmt.Println(strings.Contains("ABCDE", "CDE"))
	fmt.Println(strings.Contains("ABCDE", "CDE"))
	fmt.Println(strings.Contains("ABC", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println("---- ContainsAny ----")
	fmt.Println(strings.ContainsAny("ABC", "AE"))
	fmt.Println(strings.ContainsAny("ABC", "Cookbook"))
	fmt.Println(strings.ContainsAny("ABC", "XYZcookbook"))
}

func joinfunc() {
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))
	fmt.Println(strings.Join([]string{"Hello", ", ", "World"}, ""))
}

func indexfunc() {
	fmt.Println("---- Index ----")
	fmt.Println(strings.Index("ABCDE", "A"))
	fmt.Println(strings.Index("ABCDE", "BCD"))
	fmt.Println(strings.Index("ABCDE", "X"))
	fmt.Println(strings.Index("ABCABCABC", "ABC"))
	fmt.Println(strings.LastIndex("ABCABCABC", "ABC"))

	fmt.Println("---- IndexAny ----")
	fmt.Println(strings.IndexAny("ABC", "ABC"))
	fmt.Println(strings.IndexAny("ABC", "BCD"))
	fmt.Println(strings.IndexAny("ABC", "CDE"))
	fmt.Println(strings.IndexAny("ABC", "XYZ"))
	fmt.Println(strings.LastIndexAny("ABC", "ABC"))
	fmt.Println(strings.LastIndexAny("ABC", "XYZ"))
}
