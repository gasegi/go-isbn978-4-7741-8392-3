package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	doSomething()
	fmt.Printf("Hello, world.\n")
}

func doSomething() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "ipconfig")
	} else {
		cmd = exec.Command("/bin/sh", "-c", "ip addr")
	}
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
