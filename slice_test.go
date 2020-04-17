package main_test

import (
	"fmt"
	"testing"
)

func TestAppendSelf(t *testing.T) {
	double := func(s []int) {
		s = append(s, s...)
		//s = append(s, s[0], s[1])
	}

	s := []int{1, 2, 3}
	double(s)
	fmt.Println(s, len(s)) //prints [1,2,3] 3
}

func TestArg(t *testing.T) {
	addOne := func(s []int) []int {
		for k, v := range s {
			s[k] = v + 1
		}
		return s
	}

	s := []int{1, 2, 3}
	r := addOne(s)
	fmt.Println("origin:", s)   //[2,3,4]
	fmt.Println("returned:", r) //[2,3,4]
}

func TestCopy(t *testing.T) {

	s := []int{1, 2, 3}

	var t0 []int
	copy(t0, s)

	fmt.Println("Source:", s)
	fmt.Println("Target:", t0)

	t1 := make([]int, len(s))
	copy(t1, s)
	fmt.Println("Target:", t1)

}
