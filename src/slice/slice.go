package slice

import (
	"fmt"
)

func SliceTest() {
	var s1 []int
	var s2 = []int{0, 1, 2}
	var a1 [5]int
	s3 := a1[:1]

	fmt.Println("s1 :", s1, "s2 :", s2, "s3 :", s3)
	fmt.Println("test append change cap of slice.......")
	fmt.Println()
	if cap(s1) == 0 {
		fmt.Println("cap size : 0, append num to s1")
		s1 = append(s1, 0)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("befor appen, cap size :", cap(s1), ", len of s1", len(s1))
		ac := cap(s1) + 1 - len(s1)
		fmt.Println("append num to s1 over cap, append size :", ac)
		for j := 0; j < ac; j++ {
			s1 = append(s1, 0)
		}
		fmt.Println("after append, cap size of s1 :", cap(s1))
		fmt.Println("continue...")
	}

	fmt.Println("init slice using make")
	s4 := make([]int, 5, 10)
	fmt.Println("s4 := make([]int, 5, 10), len s4:", len(s4), ", cap s4 :", cap(s4))
	s5 := make([]int, 10)
	fmt.Println("s5 := make([]int, 10), len s5 :", len(s5), ", cap s5 :", cap(s5))
}
