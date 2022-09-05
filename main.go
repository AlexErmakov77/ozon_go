package main

import (
	"fmt"
)

func main() {
	var a, b, result int16
	var n_test, i uint16
	fmt.Scan(&n_test)
	for i = 0; i < n_test; i++ {
		fmt.Scan(&a, &b)
		result = a + b
		fmt.Println(result)
	}
}
