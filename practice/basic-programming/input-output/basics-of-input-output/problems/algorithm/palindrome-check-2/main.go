package main

import "fmt"

func main() {

	var s string
	fmt.Scan(&s)

	if s == reverse(s) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func reverse(s string) string {
	a := []rune(s)
	r := []rune{}
	for i := len(a); i > 0; i-- {
		r = append(r, a[i-1])
	}
	return string(r)
}
