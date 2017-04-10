package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bmi(a, b float64) string {
	bmi := a / (b * b)
	if bmi < 18.5 {
		return "under"
	} else if bmi >= 18.5 && bmi < 25.0 {
		return "normal"
	} else if bmi >= 25.0 && bmi < 30.0 {
		return "over"
	}
	return "obese"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')
	d := strings.Replace(num, "\n", "", -1)
	n, _ := strconv.ParseInt(d, 10, 64)

	arrData := make([]string, 0)
	var i int64

	for i = 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		data := strings.Split(input, " ")

		a, _ := strconv.ParseFloat(data[0], 64)
		b, _ := strconv.ParseFloat(data[1], 64)

		c := bmi(a, b)
		arrData = append(arrData, c)
	}
	for i := range arrData {
		fmt.Println(arrData[i])
	}
}
