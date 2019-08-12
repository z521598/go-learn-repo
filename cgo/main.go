package main

//#include "myabs.h"
//#cgo CFLAGS: -I.
import "C"
import "fmt"

func GoMyAbs(x int64) int64 {
	return int64(C.myabs(C.int(x)))
}

func main() {
	var x int64 = -128
	fmt.Println(GoMyAbs(x))
}
