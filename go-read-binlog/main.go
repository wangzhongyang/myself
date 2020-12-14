package main

import (
	"fmt"
	"reflect"
	"sync"
)

type I interface {
	Print()
}

type A struct {
	Name string
}

func (a A) Print() {
	fmt.Println(a.Name)
}

type B struct {
	Age int
}

func (b B) Print() {
	fmt.Println(b.Age)
}

func Print(i I) {
	i.Print()
}

func (b *B) IsEmpty() bool {
	return reflect.ValueOf(b).IsNil() && reflect.ValueOf(b).IsValid()
}

func AAAA() *B {
	return nil
}

func main() {
	sync.Pool{}
	//b1 := AAAA()
	//fmt.Println(b1.IsEmpty())
	////b2 := &B{Age: 3}
	////fmt.Println(b2.IsEmpty())
}
