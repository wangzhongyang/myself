package p

import "fmt"

type V struct {
	i int
	j string
	k float64
}

// 此处不用V，否则this 会是调用者的复制，而不是原对象
func (this *V) PutIPtr() {
	fmt.Println(fmt.Sprintf("v.i ptr:%p", &this.i))
}

func (this V) PutI() {
	fmt.Println(fmt.Sprintf("V.i:%d", this.i))
}

func (this V) PutJ() {
	fmt.Println(fmt.Sprintf("V.j:%s", this.j))
}

func (this V) PutK() {
	fmt.Println(fmt.Sprintf("V.j:%f", this.k))
}
