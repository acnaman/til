package main

import (
	"fmt"
	"log"
	"os"
	path "path/filepath"
)

func main() {
	
}

// call readdirfunc(0, ".")
func readdirfunc(indentNum int, dirPath string) {
	f, err := os.Open(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fis, err := f.Readdir(0)

	indent := ""
	for i := 0; i < indentNum; i++ {
		indent += "  "
	}
	
	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Println(indent + "[" + fi.Name() + "]")
			readdirfunc(indentNum + 1, path.Join(dirPath, fi.Name()))
		} else {
			fmt.Println(indent + fi.Name())
		}
	}
}

func filecreatefunc() {
	f, _ := os.Create("foo.txt")
	defer f.Close()

	fi, _ := f.Stat()
	name := fi.Name()
	size := fi.Size()
	isdir := fi.IsDir()
	fmt.Println(name, size, isdir)

	f.Write([]byte("Hello, World!\n"))
	f.WriteAt([]byte("Golang"), 7)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yeah!")
}

func filefunc() {
	var (
		bs []byte
		n int
		err error
		offset int64
		fi os.FileInfo
	)
	
	f, err := os.Open("test.txt")
	if err != nil  {
		log.Fatal(err)
	}
	defer f.Close()

	bs = make([]byte, 128)
	n, err = f.Read(bs)

	fmt.Println(n, string(bs))

	n, err = f.ReadAt(bs, 10)

	fmt.Println(n, string(bs))

	offset, err = f.Seek(10, os.SEEK_SET)
	fmt.Println(offset)
	offset, err = f.Seek(-2, os.SEEK_CUR)
	fmt.Println(offset)
	offset, err = f.Seek(0, os.SEEK_END)
	fmt.Println(offset)
	
	fi, err = f.Stat()
	name := fi.Name()
	size := fi.Size()
	time := fi.ModTime()
	isdir := fi.IsDir()

	fmt.Println(name, size, time, isdir)
}

func openfunc() {
	f, err := os.Open("test.txt")
	if err != nil  {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(f)
}

func argsfunc() {
	fmt.Printf("length=%d\n", len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
}

func exitfunc() {
	defer func() {
		fmt.Println("defer")
	}()

	os.Exit(7)
}

func logfunc() {
	_, err := os.Open("foo")
	if err != nil {
		log.Fatal(err)
	}
}
