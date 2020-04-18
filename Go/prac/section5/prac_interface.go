package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	ifunc2()
}

func ifunc1() {
	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return "hogehoge"
}

func ifunc2() {
	t := &T{Id: 1, Name: "name"}
	fmt.Println(t)
	t2 := T{Id: 1, Name: "name"}
	fmt.Println(t2.String())

	i3 := NewI3()
	res := i3.Method3()
	fmt.Println(res)
}

type I0 interface {
	Method1() int
}

type I1 interface {
	I0
	Method2() int
}

type I3 interface {
	I1
	Method3() int
}

type TI struct{}

func (t *TI) Method1() int {
	return 1
}

func (t *TI) Method2() int {
	return 1
}

func (t *TI) Method3() int {
	return 1
}

func NewI3() I3 {
	return &TI{}
}
