package main

import (
	"fmt"
	"regexp"
)

func submatchfunc() {
	re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
00-1111-2222
3333-44-55
666-777-888
9-9-9
`
	ms := re.FindAllStringSubmatch(s, -1)
	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(re.ReplaceAllString("Tell: 000-111-222", "$3-$2-$1"))
}

func regusagefunc() {
	re := regexp.MustCompile(`\s+`)

	fmt.Println(re.Split("A B C D\tE", 3))
	fmt.Println(re.Split("A B C D\tE", -1))

	re2 := regexp.MustCompile(`佐藤`)
	fmt.Println(re2.ReplaceAllString("佐藤さんと鈴木さん", "田中"))
}

func regtypefunc() {
	re1 := regexp.MustCompile(`^XYZ$`)
	s1 := "XYZ"
	s2 := "  XYZ  "
	s3 := "  XYZ  \nXYZ\n  XYZ  "
	fmt.Println(re1.MatchString(s1))
	fmt.Println(re1.MatchString(s2))
	fmt.Println(re1.MatchString(s3))

	re2 := regexp.MustCompile(`(?m)^XYZ$`)
	fmt.Println(re2.MatchString(s1))
	fmt.Println(re2.MatchString(s2))
	fmt.Println(re2.MatchString(s3))
}

func compilefunc() {
	re1 := regexp.MustCompile("A")
	fmt.Println(re1.MatchString("ABC"))
	re2 := regexp.MustCompile(`(?i)abc`)
	fmt.Println(re2.MatchString("ABC"))

}

func matchstringfunc() {
	fmt.Println(regexp.MatchString("A", "ABC"))
	fmt.Println(regexp.MatchString("A", "XYZ"))
}
