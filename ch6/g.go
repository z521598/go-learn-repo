package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
)

func Say(word string, t time.Duration, dataCh chan string, errorCh chan error) {
	if word == "hello" {
		dataCh <- uuid.New().String() + " end"
	} else if word == "go" {
		errorCh <- fmt.Errorf("not hello")
	} else {
		runtime.Goexit()
	}
	for i := 0; i < 5; i++ {
		//time.Sleep(t)
		fmt.Println(word)
	}

}

func main() {
	runtime.GOMAXPROCS(1)
	c1 := make(chan string)
	c11 := make(chan error)
	c2 := make(chan string)
	c22 := make(chan error)
	go Say("hello", 1*time.Second, c1, c11)
	go Say("go", 2*time.Second, c2, c22)
	go Say("00", 2*time.Second, nil, nil)

	select {
	case res1 := <-c1:
		fmt.Println(res1)
	case err := <-c11:
		fmt.Println(err)
	default:
		fmt.Println("未结束")
	}

	select {
	case res2 := <-c2:
		fmt.Println(res2)
	case err := <-c22:
		fmt.Println(err)
	default:
		fmt.Println("未结束")
	}
}
