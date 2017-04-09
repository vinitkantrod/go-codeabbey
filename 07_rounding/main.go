package main

import (
	"bufio"
	"fmt"
	"math"
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
	result := make([]float64, 0)
	for i = 0; i < n; i++ {
		data, _ := reader.ReadString('\n')
		data = strings.Replace(data, "\n", "", -1)
		values := strings.Split(data, " ")

		a, _ := strconv.ParseFloat(values[0], 64)
		b, _ := strconv.ParseFloat(values[1], 64)

		c := (a / b)
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
		result = append(result, c)
	}
	for i := range result {
		fmt.Println(result[i])
	}
}
