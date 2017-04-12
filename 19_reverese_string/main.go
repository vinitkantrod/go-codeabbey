package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	rune := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}
	return string(rune)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	res := reverse(input)
	fmt.Println(res)
}
