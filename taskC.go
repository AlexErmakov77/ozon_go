package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nInput, nUnit, i, j int

	fmt.Fscan(in, &nInput)

	for j = 0; j < nInput; j++ {
		fmt.Fscan(in, &nUnit)
		allUnit := make([]int, nUnit, nUnit)

		for i = 0; i < nUnit; i++ {
			fmt.Fscan(in, &allUnit[i])
		}

		var count1 int

		for i = 0; i < nUnit/2; i++ {

			count1 = 0
			for {
				if allUnit[count1] != 0 {
					break
				}
				count1++
			}

			var count2, min_dif, k int
			min_dif = 100

			for k = count1 + 1; k < nUnit; k++ {

				if allUnit[k] != 0 && (min_dif > int(math.Abs(float64(allUnit[count1])-float64(allUnit[k])))) {
					count2 = k
					min_dif = int(math.Abs(float64(allUnit[count1]) - float64(allUnit[k])))
				}
			}
			fmt.Println(count1+1, count2+1)
			allUnit[count1] = 0
			allUnit[count2] = 0
		}
		fmt.Println()
	}
}
