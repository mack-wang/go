package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[3:7]
	s2 := a[1:4]

	// 影响了a,s1,s2
	for i := range s2 {
		s2[i] += 100
	}

	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)

}
