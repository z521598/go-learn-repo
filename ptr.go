package main
import "fmt"

type People struct {
	age int
	name string
}

func change(p People){
	fmt.Println("change");
	fmt.Println(&p);
	p.age = 18
}
func change2(p *People){
	fmt.Println("change2");
	fmt.Println(&p);
	p.age = 18
}

func change3(p *People){
	fmt.Println("change3");
	fmt.Println(&p);
	p = new(People)
	p.age = 29
	p.name= "liming07"
}
func main(){
        p := People{22,"lsq"}
	fmt.Println("before change");
	fmt.Println(&p);
        change(p)
        change2(&p)
	change3(&p)
}
