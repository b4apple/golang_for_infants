package basics

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