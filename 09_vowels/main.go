package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')
	d := strings.Replace(num, "\n", "", -1)
	n, _ := strconv.ParseInt(d, 10, 64)

	arrData := make([]int64, 0)
	var i int64
	for i = 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		var vowelCount int64
		vowelCount = 0
		for i := range input {
			if 'a' == input[i] || 'e' == input[i] || 'i' == input[i] || 'o' == input[i] || 'u' == input[i] || 'y' == input[i] {
				vowelCount = vowelCount + 1
			}
		}
		arrData = append(arrData, vowelCount)
	}
	for i := range arrData {
		fmt.Println(arrData[i])
	}
}
