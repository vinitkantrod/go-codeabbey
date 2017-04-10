package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func weightedSumOfDigits(inputData string) int64 {
	var i int64
	l := int64(len(inputData))
	var sum int64
	for i = 1; i < l+1; i++ {
		ii := string(inputData[i-1])
		char, _ := strconv.ParseFloat(ii, 10)
		sum = sum + (int64(char) * i)
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')
	num = strings.Replace(num, "\n", "", -1)

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	data := strings.Split(input, " ")

	var i int64
	resData := make([]int64, 0)

	l := len(data)
	for i = 0; i < int64(l); i++ {
		c := weightedSumOfDigits(data[i])
		resData = append(resData, c)
	}
	for i := range resData {
		fmt.Println(resData[i])
	}
}
