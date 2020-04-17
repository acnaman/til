package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	func10()
}

type (
	MyInt       int
	IntPair     [2]int
	Strings     []string
	AreaMap     map[string][2]float64
	IntsChannel chan []int
)

func func1() {

	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1, n2)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)

	go func() {
		for {
			ints := <-ich
			fmt.Println("chan test: ", ints)
		}
	}()

	ich <- []int{1, 3, 45, 21}

	fmt.Println(pair, strs, amap, ich)

}

type Point struct {
	X, Y int
}

func func2() {
	var p1 Point
	p1.X = 2
	p1.Y = 3

	fmt.Println(p1)

	p2 := Point{X: 7, Y: 9}
	fmt.Println(p2)

}

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed
}

func func3() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a)
	fmt.Println(a.Amount)
}

func func4() {
	p := Point{X: 1, Y: 2}
	swap(p)
	fmt.Println(p)

	swap2(&p)
	fmt.Println(p)

	pt := new(Point)
	fmt.Printf("%T\n", pt)

	pt2 := &Point{X: 1, Y: 4}
	fmt.Printf("%T\n", pt2)

	pt2.Render()

	pt3 := NewPoint(56, 78)
	fmt.Println(pt3)

}

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

func NewPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

func swap(p Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func swap2(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func func5() {
	f := (*Point).ToString
	str := f(NewPoint(32, 56))
	fmt.Println(str)
}

func func6() {
	ps := make([]Point, 5)
	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}
}

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d,%d]", p.X, p.Y)
		}
	}
	return str
}

func func7() {
	ps := Points{}
	ps = append(ps, &Point{X: 1, Y: 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{X: 3, Y: 4})

	fmt.Println(ps.ToString())
}

type User struct {
	Id   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  uint   `json:"user_age"`
}

func func8() {
	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "Tokyo",
		{Id: 2, Name: "Jiro"}: "Osaka",
	}

	m2 := map[int]User{
		1: {Id: 1, Name: "Taro"},
		2: {Id: 2, Name: "Jiro"},
	}

	fmt.Println(m1)
	fmt.Println(m2)
}

func func9() {
	u := User{Id: 1, Name: "Taro", Age: 32}

	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}

type User1 struct {
	Id   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  uint   `json:"user_age"`
}

type User2 struct {
	Id   int
	Name string
	Age  uint
}

func func10() {
	u1 := User1{Id: 1, Name: "Taro", Age: 32}
	u2 := User2{Id: 1, Name: "Taro", Age: 32}
	bs1, _ := json.Marshal(u1)
	bs2, _ := json.Marshal(u2)
	fmt.Println(string(bs1))
	fmt.Println(string(bs2))
}
