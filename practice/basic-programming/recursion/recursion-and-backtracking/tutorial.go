package main

import (
	"fmt"
	"os"
)

var upperlim = 1
var sum = 0
var n = 4

var a []int

func main() {
	fmt.Scanf("%d", &n)
	if e := judgeNum(n); e != nil {
		return
	}
	upperlim = (upperlim << uint(n)) - 1
	test(0, 0, 0)
	if sum == 0 {
		fmt.Println("Not possible")
	}
}

func test(row, ld, rd int) {
	if row != upperlim {
		pos := upperlim & ^(row | ld | rd)
		for pos != 0 {
			p := pos & -pos
			pos -= p
			a = append(a, p)
			test(row+p, (ld+p)<<1, (rd+p)>>1)
			a = a[:len(a)-1]
		}
	} else {
		sum++
		print()
		os.Exit(0)
	}
}

func print() {
	for i := 0; i < len(a); i++ {
		num := getCount(a[i])
		for j := 0; j < n; j++ {
			if j == num {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
			if j != n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getCount(n int) (count int) {
	for n != 1 {
		n = n >> 1
		count++
	}
	return
}

func judgeNum(i int) error {
	if i < 1 || i > 10 {
		return fmt.Errorf("%d is out of 1 to 10", i)
	}
	return nil
}
