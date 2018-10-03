package main
import "fmt"

type People struct {
	id int
	name string
}

func main(){
	p1 := People{1,"lsq"}
	change(p1)
	fmt.Println(p1)
	change2(&p1)	
	fmt.Println(p1)
}
func change(p People){
	p.id = 12
}
func change2(p *People){
	p.id = 12
}
