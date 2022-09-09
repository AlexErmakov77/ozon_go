package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const MAXSEC = 24 * 60 * 60

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nTest, nTask, i, n int

	fmt.Fscan(in, &nTest)

	for n = 0; n < nTest; n++ {

		fmt.Fscan(in, &nTask)
		flag := true
		var h1, m1, s1, h2, m2, s2, t1, t2 int
		sec := make([]int, MAXSEC, MAXSEC)

		for i = 0; i < nTask; i++ {
			in1 := bufio.NewReader(os.Stdin)

			fmt.Fscanf(in1, "%d:%d:%d-%d:%d:%d", &h1, &m1, &s1, &h2, &m2, &s2)

			if checkTime(h1, m1, s1) == false || checkTime(h1, m1, s1) == false {
				flag = false
			} else {
				t1 = 3600*h1 + 60*m1 + s1
				t2 = 3600*h2 + 60*m2 + s2

				if t1 > t2 {
					flag = false
				} else {
					for j := t1; j <= t2; j++ {
						if sec[j] == 1 {
							flag = false
							break
						} else {
							sec[j] = 1
						}
					}
				}
			}
		}

		if flag == true {
			fmt.Println("YES")
		} else if flag == false {
			fmt.Println("NO")
		}
	}
}

func checkTime(h, m, s int) bool {
	var res bool
	if h >= 0 && h <= 23 && m >= 0 && m <= 59 && s >= 0 && s <= 59 {
		res = true
	} else {
		res = false
	}
	return res
}
