package other_dir

import "fmt"

type A struct {
	Name string
}

var (
	_A *A
	B  *A
)

func init() {
	_A = &A{Name: "000000"}
	B = _A
}

func Print() {
	fmt.Println("a name:", _A.Name)
}
