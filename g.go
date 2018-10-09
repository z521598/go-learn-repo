package main

import(
	"fmt"
	"runtime"
)

func say(word string){
	for i:=0 ; i < 10 ; i++ {
//		runtime.Gosched()
		fmt.Println(word);
	}
}

func main(){
	runtime.GOMAXPROCS(1)
	go say("hello")
	runtime.Gosched()
	say("go")

}
