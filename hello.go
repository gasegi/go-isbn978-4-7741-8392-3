package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	doSomething()
	fmt.Printf("Hello, world.\n")
}

func doSomething() {
	cmd := exec.Command("ipconfig")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// バイト配列 b は CP932 かもしれない
	fmt.Print(string(b))
}
