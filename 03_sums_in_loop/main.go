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

		arrData := make([]int64, 0)
		var i int64
		l := int64(len(values))
		for i = 0; i < l; i++ {
			temp, _ := strconv.ParseInt(values[i], 10, 64)
			arrData = append(arrData, temp)
		}
		var sum int64
		for _, val := range arrData {
			sum = sum + val
		}
		finalArr = append(finalArr, sum)
	}
	for i := range finalArr {
		fmt.Println(finalArr[i])
	}
}
