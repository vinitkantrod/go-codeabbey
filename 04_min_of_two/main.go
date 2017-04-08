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
	// Read number of elements in array
	num, _ := reader.ReadString('\n')
	d := strings.Replace(num, "\n", "", -1)
	n, _ := strconv.ParseInt(d, 10, 64)

	var i int64
	finalArr := make([]int64, 0)

	for i = 0; i < n; i++ {
		data, _ := reader.ReadString('\n')
		data = strings.Replace(data, "\n", "", -1)
		values := strings.Split(data, " ")
		a, _ := strconv.ParseInt(values[0], 10, 64)
		b, _ := strconv.ParseInt(values[1], 10, 64)
		if a < b {
			finalArr = append(finalArr, a)
		} else {
			finalArr = append(finalArr, b)
		}
	}
	for i := range finalArr {
		fmt.Println(finalArr[i])
	}
}
