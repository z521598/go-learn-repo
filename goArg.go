package main

import(
	"fmt"
	"os"
	"github.com/docker/docker/pkg/term"
)

func main(){
	fmt.Println(os.Args[1:])
	fmt.Println("hello,world")
	stdin, stdout, stderr := term.StdStreams()
	fmt.Println(stdin)
	fmt.Println(stdout)
	fmt.Println(stderr)
}
