package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	doSomething()
	fmt.Printf("Hello, world.\n")
}

func doSomething() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	logrus.Info("succeeded")
	logrus.Warn("not correct")
	logrus.Error("something error")
	logrus.Fatal("panic")
}
