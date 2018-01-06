package main

import (
	"strings"
	"fmt"
)

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main()  {
	fmt.Println(basename2("/a/b/cc.go"))
}
