package main

import (
	"log"
	"os"
)

// func main() {
// logmain()
// }
func logmain() {
	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	defer f.Close()
	log.SetOutput(f)

	loggerfunc()
}

func loggerfunc() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatalln("message")
}

func logfunc() {
	log.Println("test1")
	log.Println("test2")
	log.Println("test3")
	log.Panicln("kernel panic!!!")
}
