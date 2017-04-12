package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arithmaticProgression(a, b, c int64) int64 {
	if c <= 0 {
		return a
	}
	var i int64
	var sum int64
	for i = 0; i < c; i++ {
		sum = sum + (a + (i * b))
	}
	return sum
}

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
		data := strings.Split(input, " ")

		a, _ := strconv.ParseInt(data[0], 10, 64)
		b, _ := strconv.ParseInt(data[1], 10, 64)
		c, _ := strconv.ParseInt(data[2], 10, 64)
		res := arithmaticProgression(a, b, c)
		arrData = append(arrData, res)
	}
	for i := range arrData {
		fmt.Println(arrData[i])
	}
}
