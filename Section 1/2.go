package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	myslice := []int{1, 2, 3, 4, 5} //[]int{}
	fmt.Println(a)
	myslice = append(myslice, 6, 7, 8)
	myslice = append(myslice, 9)
	fmt.Println(myslice)

	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 4, 5 //capacity: cap(s) ,length: len(s)
	fmt.Println(s)
	s1 := s[1:3]
	fmt.Println(s1[:cap(s1)])
	fmt.Println(gettwo(s, 1, 3))

}

func gettwo(s []int, fi, si int) []int { // if s has 100 elements, then the return value will have access to 100-fi elements
	s2 := make([]int, si-fi)
	copy(s2, s[fi:si])
	return s2
}
