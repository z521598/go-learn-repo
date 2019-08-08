package main

import (
	"os"
	"time"
)

func main() {
	c1 := make(chan string)
	c11 := make(chan error)
	go Say("123", 1*time.Second, c1, c11)

}
