package main

import (
	"fmt"
	"time"
)

func tickfunc() {
	ch := time.Tick(3 * time.Second)

	for {
		t := <-ch
		fmt.Println(t)
	}
}

func adddatefunc() {
	t := time.Date(1989, 9, 19, 0, 0, 0, 0, time.Local)

	fmt.Println(t.Before(time.Now()), t.Format("2006/1/2"))

	t = t.AddDate(1000, 0, 0)

	fmt.Println(t.Before(time.Now()), t.Format("2006/01/02"))
}

func timeweekdayfunc() {
	fmt.Println(time.July.String())
}

func timefunc() {
	t := time.Now()

	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Zone())

}
