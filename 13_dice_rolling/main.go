package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const factor = 6.0

func diceValue(a float64) int64 {
	factorRes := a * factor
	return int64(math.Floor(factorRes) + 1.0)
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
		a, _ := strconv.ParseFloat(input, 64)
		c := diceValue(a)
		arrData = append(arrData, c)
	}
	for i := range arrData {
		fmt.Println(arrData[i])
	}
}
