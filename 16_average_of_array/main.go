package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func round(c float64) int64 {
	intC := int64(c / 0.5)
	if intC%2 == 0 {
		if c < 0 {
			c = math.Ceil(c)
		} else {
			c = math.Floor(c)
		}
	} else {
		if c < 0 {
			c = math.Floor(c)
		} else {
			c = math.Ceil(c)
		}
	}
	if c == 0 || c == -0 {
		c = 0
	}
	return int64(c)
}

func getAverageOfArray(newData []float64) int64 {
	var sum float64
	for i := range newData {
		sum = sum + newData[i]
	}
	c := sum / float64(len(newData))
	return round(c)
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
		newData := make([]float64, 0)
		for i := range data {
			temp, _ := strconv.ParseFloat(data[i], 64)
			newData = append(newData, temp)
		}
		newData = newData[:len(newData)-1]
		c := getAverageOfArray(newData)
		arrData = append(arrData, c)
	}
	for i := range arrData {
		fmt.Println(arrData[i])
	}
}
