package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Returns sum of digits for an integer
func sumOfDigits(a, b, c int64) int64 {
	s := (a * b) + c
	ss := strconv.FormatInt(s, 10) // formatInt take int64 and string base (10) and return string
	l := len(ss)
	var i int
	var sum int64
	// each charter is uint8 and we converted uint8 into string into float64 and then into int64
	for i = 0; i < l; i++ {
		ii := string(ss[i])
		t, _ := strconv.ParseFloat(ii, 10)
		sum = sum + int64(t)
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
		res := sumOfDigits(a, b, c)
		arrData = append(arrData, res)
	}
	for i := range arrData {
		fmt.Println(arrData[i])
	}
}
