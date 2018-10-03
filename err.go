package main

import(
	"fmt"
//	"errors"
)

func exec() error {
//	return errors.New("err msg") 
	
	return nil
}

func main() {
	if error := exec();error != nil {
		fmt.Println(error)
	}
}
