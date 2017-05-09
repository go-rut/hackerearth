package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	bs, e := inputReader.ReadBytes('\n')
	if e != nil && e != io.EOF {
		fmt.Println("input string failed1:", e)
		return
	}
	lenS := len(bs)

	if lenS < 1 || lenS > 100 {
		fmt.Println("length of string is out of 1 to 100")
		return
	}

	counts := make(map[byte]int, 10)

	for i := 0; i < lenS; i++ {
		if bs[i] < 48 || bs[i] > 57 {
			continue
		}
		counts[bs[i]]++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, counts[[]byte(strconv.Itoa(i))[0]])
	}
}
