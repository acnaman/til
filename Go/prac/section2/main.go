package main

import (
	"fmt"

	"./animals"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(animals.MonkeyFeed())
}
