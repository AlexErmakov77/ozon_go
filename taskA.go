package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n_test, i uint16
	fmt.Fscan(in, &n_test)

	for i = 0; i < n_test; i++ {
		var a, b, result int16
		fmt.Fscan(in, &a, &b)
		result = a + b
		fmt.Fprintln(out, result)
	}
}
