package intermediate

import "fmt"

//Iter1 return the factorial of 5
func Iter1() {
	var factorial int64 = 1
	for i := 1; i <= 5; i++ {
		factorial *= int64(i)
	}
	fmt.Printf("Factorial of 5 is %v", factorial)
}
