package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/text/encoding/japanese"
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
	// コマンド実行からの出力は CP932 であると限定し UTF-8 へ変換する
	b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}
