package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	calendar := []int{29, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n_test, i int
	fmt.Fscan(in, &n_test)

	result := make([]string, n_test, n_test)

	for i = 0; i < n_test; i++ {
		var d, m, y int
		var res bool

		fmt.Fscan(in, &d, &m, &y)

		if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
			m = 0
		}

		if d <= calendar[m] {
			res = true
		} else {
			res = false
		}

		if res {
			result[i] = "YES"
			//	fmt.Println("YES")
		} else {
			result[i] = "NO"
			//	fmt.Println("NO")
		}
	}

	for i = 0; i < n_test; i++ {
		fmt.Println(result[i])
	}

}
