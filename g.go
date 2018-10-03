package main

import(
	"fmt"
	_ "runtime"
)

func say(word string){
	for i:=0 ; i < 10 ; i++ {
//		runtime.Gosched()
		fmt.Println(word);
	}
}

func main(){
	go say("hello")
	say("go")

}
