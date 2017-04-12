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

	num, _ := reader.ReadString('\n')
	d := strings.Replace(num, "\n", "", -1)
	_, er := strconv.ParseInt(d, 10, 64)
	if er != nil {
		fmt.Println("Error")
	}

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	data := strings.Split(input, " ")

	var result int64
	for i := range data {
		temp, _ := strconv.ParseInt(data[i], 10, 64)
		result = result + temp
		result = (result * 113) % 10000007
	}
	fmt.Println(result)
}
