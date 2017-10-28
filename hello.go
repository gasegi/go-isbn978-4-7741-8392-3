package main

import (
	"fmt"
	"os"
)

func main() {
	doSomething()
	fmt.Printf("Hello, world.\n")
}

func doSomething() error {
	err := os.MkdirAll("newdir", 0755)
	if err != nil {
		return err
	}
	// (2) 次にディレクトリが削除される
	defer os.RemoveAll("newdir")

	f, err := os.Create("newdir/newfile")
	if err != nil {
		return err
	}
	// (1) 最初にファイルハンドルが閉じられる
	defer f.Close()
	return nil
}
