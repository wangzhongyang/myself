package test2

import (
	"fmt"
	"testing"
)

func TestPrint2(t *testing.T) {
	t.Run("this is test 2", func(t *testing.T) {
		fmt.Println("this is test 2")
	})
}
