/*
	f °Fahrenheit to c °Celsius : ((f - 32) * 5) / 9 = c
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func round(c float64) float64 {
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
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Read number of elements in array
	data, _ := reader.ReadString('\n')
	data = strings.Replace(data, "\n", "", -1)
	values := strings.Split(data, " ")

	arrData := make([]float64, 0)
	var i int64
	l := int64(len(values))
	for i = 0; i < l; i++ {
		temp, _ := strconv.ParseFloat(values[i], 64)
		arrData = append(arrData, temp)
	}
	calValues := arrData[1:]
	result := make([]float64, 0)
	for i := range calValues {
		c := ((calValues[i] - 32) * 5) / 9
		result = append(result, c)
	}
	for i := range result {
		c := round(result[i])
		fmt.Println(c)
	}
}
