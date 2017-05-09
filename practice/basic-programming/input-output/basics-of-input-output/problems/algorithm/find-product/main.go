package main

import (
	"fmt"
)

func main() {
	var n int
	var j int
	result := 1
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&j)
		result = (result * j) % 1000000007
	}
	fmt.Println(result)
}
