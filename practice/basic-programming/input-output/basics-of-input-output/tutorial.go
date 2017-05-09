package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var i int
	inputReader := bufio.NewReader(os.Stdin)
	s, e := inputReader.ReadString('\n')
	if e != nil {
		fmt.Println("input string failed1:", e)
		return
	}
	_, e = fmt.Sscanf(s, "%d\n", &i)
	if e != nil {
		fmt.Println("input string failed2:", e)
		return
	}
	if i < 0 || i > 10 {
		fmt.Println("length of int is out 1 to 10")
		return
	}

	s, e = inputReader.ReadString('\n')
	if e != nil && e != io.EOF {
		fmt.Println("input string failed3:", e)
		return
	}
	lenS := len(s)
	if lenS < 1 || lenS > 15 {
		fmt.Println("length of string is out 1 to 15")
		return
	}

	fmt.Println(i * 2)
	fmt.Println(s)
}
