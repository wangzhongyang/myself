package test1

import (
	"fmt"
	"testing"
)

func TestPrint1(t *testing.T) {
	t.Run("this is test 1", func(t *testing.T) {
		fmt.Println("this is test 1")
	})
}

func TestPrint3(t *testing.T) {
	t.Run("this is test 3", func(t *testing.T) {
		fmt.Println("this is test 3")
		t.Fatal("this is test fatal")
	})
}
