package main_test

import (
	"fmt"
	"testing"
)

func TestSliceArg(t *testing.T) {
	double := func(s []int) {
		s = append(s, s...)
	}

	s := []int{1, 2, 3}
	double(s)
	fmt.Println(s, len(s)) //prints [1,2,3] 3
}
