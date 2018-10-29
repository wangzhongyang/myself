package main

type People interface {
	Run()
}
type Young struct {
	Name int
	Age  int
}

func (y Young) Run() {
	println("name:%d,age:%d", y.Name, y.Age)
}

type Old struct {
	Name int
}

func (o Old) Run() {
	println("i am only have name", o.Name)
}

func Run(people People) {
	people.Run()
}

func main() {
	y := Young{Name: 1, Age: 2}
	o := Old{Name: 3}
	Run(y)
	Run(o)
}
