package main

import (
	"fmt"
	"strings"
)

func toggleCase(s string) string {
	ss := strings.ToUpper(s)
	if ss == s {
		return strings.ToLower(s)
	}
	return ss
}

func main() {

	var s string

	fmt.Scanf("%s", &s)
	var rs string
	for _, v := range s {
		rs += toggleCase(string(v))
	}
	fmt.Println(rs)
}
