package main

import "fmt"
import "errors"

func main() {
	fmt.Println("hello,world")
	err := errors.New("my error")
	if err != nil {
		fmt.Println(err)
	}
	var i int = 1
	fmt.Println(incr(i))
	fmt.Println(i)
	j := 1
	fmt.Println(incr2(&j))
	fmt.Println(j)
}

func incr(number int) int {
	number = number + 1
	return number
}

func incr2(number *int) int {
	*number = *number + 1
	return *number
}
