//Add up the even fibonacci numbers less than 4 million
package main

import (
	"fmt"
)

func main() {
	var acc int
	x := 0
	y := 1
	var z int
	for y <= 4000000 {
		z = x + y
		x = y
		y = z
		if z%2 == 0 {
			acc += z
		}
	}
	fmt.Println(acc)
}
