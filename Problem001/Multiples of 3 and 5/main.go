//Find the sum of all of the multiples of 3 and 5 less than 1000
package main

import (
	"fmt"
)

func main() {
	var acc int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			acc += i
		}
	}
	fmt.Println(acc)
}
