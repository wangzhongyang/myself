package main

import (
	"fmt"
	"testing"
)

type Name struct {
	NameLen string `json:"name_len"`
}

func TestName(t *testing.T) {
	n := new(Name)
	n.NameLen = "111"
	fmt.Printf("\nn:        %+v\n", n)
}
