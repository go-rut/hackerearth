package main

import "fmt"

func main() {
	var t int
	var n int
	fmt.Scanf("%d", &t)
	if e := judgeNum("t", t); e != nil {
		fmt.Println(e)
		return
	}
	fmt.Scanf("%d", &n)

	if e := judgeNum("n", n); e != nil {
		fmt.Println(e)
		return
	}

	for i := 0; i < t; i++ {
		fmt.Println(getNumber(i + 1))
	}
}

func judgeNum(key string, i int) error {
	if i < 1 || i > 1000 {
		return fmt.Errorf("%s:%d is out of 1 to 1000", key, i)
	}
	return nil
}

func getNumber(n int) (count int) {
	for n > 0 {
		if n%2 != 0 {
			count++
		}
		n = n >> 1
	}

	return
}
