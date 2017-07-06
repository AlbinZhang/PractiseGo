package main

import (
	"fmt"
)

func reverse(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	for ; n > 0; n-- {
		t := s[0]
		copy(s[0:], s[1:])
		s[len(s)-1] = t
	}
}

func delRepeatStr(s []string) []string {
	var n int
	for i, v := range s {
		if i+1 == len(s) {
			break
		}

		if v != s[i+1] {
			s[n] = v
			n++
		}
	}

	return s[:n]
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6}
	reverse(&arr)
	fmt.Println(arr)

	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 4)
	fmt.Println(s)

	strSlice := []string{"111", "234", "234", "23456", "234", "222", "222"}
	fmt.Println(strSlice)
	strSlice = delRepeatStr(strSlice)
	fmt.Println(strSlice)
}
