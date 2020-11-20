package intermediate

//Iter2 returns the factorial of a number accepted from STDIN (iterative using loop)
func Iter2(number int64) int64 {
	var factorial int64 = 1
	for i := int64(1); i <= number; i++ {
		factorial *= i
	}
	return factorial
}
