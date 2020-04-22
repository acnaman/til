package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	shafunc()
}

func shafunc() {
	str := "ABCDE"

	/* sha-1 */
	s1 := sha1.New()
	io.WriteString(s1, str)
	fmt.Println(fmt.Printf("%x", s1.Sum(nil)))

	/* sha-256 */
	s256 := sha256.New()
	io.WriteString(s256, str)
	fmt.Println(fmt.Printf("%x", s256.Sum(nil)))

	/* sha-1 */
	s512 := sha512.New()
	io.WriteString(s512, str)
	fmt.Println(fmt.Printf("%x", s512.Sum(nil)))
}

func md5func() {
	h := md5.New()
	io.WriteString(h, "ABCDE")

	fmt.Println(h.Sum((nil)))

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}
