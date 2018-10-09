package main

import(
	"fmt"
	_ "context"
)

func main(){
	ch := make(chan int)
	go func(){
		x := <- ch
		fmt.Println(x)	
	}()
	ch <- 1
	fmt.Println(2)
	
}
