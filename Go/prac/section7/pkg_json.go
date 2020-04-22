package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	unmarshalfunc()
}

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func unmarshalfunc() {
	src := `
{
	"Id":12,
	"Name":"田中花子",
	"Email":"tanaka@example.com",
	"Created":"2020-04-22T09:37:46.9069412Z"
}`
	u := new(User)
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
}

func marshalfunc() {
	u := new(User)
	u.Id = 1
	u.Name = "山田太郎"
	u.Email = "yamada@example.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
