
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t) // Read the number of test cases

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n) // Read the integer n for each test case

		ans := 1
		for n > 3 {
			n /= 4
			ans *= 2
		}

		fmt.Println(ans) // Print the result for each test case
	}
}
