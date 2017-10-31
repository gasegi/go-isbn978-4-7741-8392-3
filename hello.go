package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Printf("Hello, world.\n")
}

func doSomething() {
	var b1 byte = 'a' // => '97'
	// var b2 byte = byte('あ') // ERROR
	// var b3 byte = 'あ' // ERROR

	fmt.Println(b1)
}
