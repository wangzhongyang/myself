package main

import "fmt"

func main() {
	b := get()
	b.PrintString()
}

func get() A {
	return NewB()
}

type A interface {
	PrintString()
}
type B string

func NewB() *B {

	return new(B)
}

func (b *B) PrintString() {
	fmt.Println("bbbbb")
}

func (b *B) PrintInt() {
	fmt.Println("int ")
}
