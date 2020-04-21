package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	tempfilefunc()
}

func tempfilefunc() {
	f, err := ioutil.TempFile(os.TempDir(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("Hello, World!")

	fmt.Println(f.Name())
}

func readallfunc() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}

func flushfunc() {
	w := bufio.NewWriter(os.Stdout)
	w.WriteString("Hello, World!\n")
	//w.Flush()
}

func scannersplitfunc() {
	s := `ABC DEF
GHI JKL MNO
PQR STU VWX
YZ`
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func stringsreaderfunc() {
	s := `XXXXX
YYYYY
ZZZZZ`
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
}

func scannerfunc() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー", err)
	}
}
