package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number = ")
	var input, _ = reader.ReadString('\n')
	var number, _ = strconv.ParseFloat(strings.Trim(input, "\n\r "), 64)
	fmt.Printf("The number entered is %v\n", number)
}
