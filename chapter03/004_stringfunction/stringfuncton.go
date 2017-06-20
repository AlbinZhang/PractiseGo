package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("abc"))
	fmt.Println(basename2("a/b/c"))
	fmt.Println(basename2("a/b.c.go"))

	fmt.Println(comma("1234567"))
	fmt.Println(commaf("1234567.89"))
}

// basename removes directory components and a .suffix
// eg., a ==> a, a/b/c.go ==> c, a/b.c.go ==> b.c
func basename1(s string) string {
	//Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//Preserve everything before last '.' .
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaf(s string) string {
	if dot := strings.LastIndex(s, "."); dot > 0 {
		return comma(s[:dot-1]) + s[dot:]
	}

	return comma(s)
}
