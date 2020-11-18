package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name")
	var givenName, _ = reader.ReadString('\n')
	fmt.Println("Your name is probabaly " + givenName + "but mine is Skaadin")
}
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
package main

import "fmt"

func variables() {
	var myLuckyNumber int64 = 100
	myName := "Skaadin"
	fmt.Println(myName)
	fmt.Println(myLuckyNumber)
}
<h1>GoLang for Infants </h1>
<p>Basically a collection of .go files that act as a cheatsheet. </p><br>
<p>This also documents my learning journey of GoLang.</p>
