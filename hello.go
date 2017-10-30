package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	doSomething()
	fmt.Printf("Hello, world.\n")
}

func doSomething() {
	f, err := os.Create("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("Hello"))

	f, err = os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write([]byte("World"))
}
