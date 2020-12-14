package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type People struct {
	Name string    `json:"name"`
	A    time.Time `json:"a"`
}

type Peoples []People

func (p Peoples) Len() int           { return len(p) }
func (p Peoples) Less(i, j int) bool { return p[i].A.Unix() < p[j].A.Unix() }
func (p Peoples) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

}

func test1() {
	a := []int{1, 2, 3, 4, 5, 6}
	a1 := make([]int, 3)
	copy(a1, a[:3])
	fmt.Println(a1)
	a[0], a[5] = a[5], a[0]
	fmt.Println(a1)
}

func return1(n int) []*People {
	p := make([]*People, 0)
	for i := 0; i < n; i++ {
		t := People{Name: fmt.Sprintf("name%d", n)}
		//fmt.Println(fmt.Sprintf("t1:%p", &t))
		p = append(p, &t)
	}
	return p
}

func return2(n int) []People {
	p := make([]People, 0)
	for i := 0; i < n; i++ {
		t := People{Name: fmt.Sprintf("name%d", n)}
		//fmt.Println(fmt.Sprintf("t2:%p", &t))
		p = append(p, t)
	}
	//fmt.Println(fmt.Sprintf("t5:%p", &p))
	return p
}

func return3(n int) *[]People {
	p := make([]People, 0)
	for i := 0; i < n; i++ {
		t := People{Name: fmt.Sprintf("name%d", n)}
		p = append(p, t)
	}
	return &p
}
