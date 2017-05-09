package main

import "fmt"

func main() {

	var l, r, k, count int
	fmt.Scan(&l)
	fmt.Scan(&r)
	fmt.Scan(&k)

	for l <= r {
		if l%k == 0 {
			l += k
			count++
		} else {
			l++
		}
	}
	fmt.Println(count)
}
