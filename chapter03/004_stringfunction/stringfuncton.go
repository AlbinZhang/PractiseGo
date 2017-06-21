package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename2("abc"))
	fmt.Println(basename2("a/b/c"))
	fmt.Println(basename2("a/b.c.go"))

	fmt.Println(commafor("123"))
	fmt.Println(commafor("1234"))
	fmt.Println(commafor("12345"))
	fmt.Println(commafor("123456"))
	fmt.Println(commafor("1234567"))
	fmt.Println(comma("1234567.89"))
	fmt.Println(intsToString([]int{1, 2, 3, 4, 5}))
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

	if dot := strings.LastIndex(s, "."); dot > 0 {
		return comma(s[:dot-1]) + s[dot:]
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func commafor(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	count := len(s) / 3

	fmt.Fprintf(&buf, "%s", s[:n])
	for i := 0; i < count; i++ {
		if i == 0 && n == 0 {
			fmt.Fprintf(&buf, "%s", s[n+i*3:n+(i+1)*3])
			continue
		}
		fmt.Fprintf(&buf, ",%s", s[n+i*3:n+(i+1)*3])
	}

	return buf.String()
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
