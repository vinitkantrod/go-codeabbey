package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	values := strings.Split(text, " ")
	// convert strings into respective Int using strconv.ParseInt
	// ParseInt take 3 args:
	// 1. String, 2. base (0 - 36), 3. bitSize (8, 16, 32, 64)==(int8, int16, int32, int64)
	// return int64 and err
	a, _ := strconv.ParseInt(values[0], 10, 64)
	b, _ := strconv.ParseInt(values[1], 10, 64)
	sum := a + b
	fmt.Println(sum)

}
