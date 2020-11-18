package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conditionals() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number : ")
	var number, _ = reader.ReadString('\n')
	var actNumber, _ = strconv.ParseInt(strings.Trim(number, "\n\r "), 10, 64)
	if actNumber%7 == 0 || actNumber%10 == 7 {
		fmt.Printf("%v is a buzz number", actNumber)
		//return
	} else {
		fmt.Printf("%v is not a buzz number", actNumber)
	}
}
