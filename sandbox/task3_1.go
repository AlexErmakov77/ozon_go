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

	var nInput, nUnit, i, j, n int

	fmt.Fscan(in, &nInput)

	sortUnit := []int{}

	for n = 0; n < nInput; n++ {

		fmt.Fscan(in, &nUnit)

		allUnit := make([]int, nUnit, nUnit)

		for i = 0; i < nUnit; i++ {
			var kval int
			fmt.Fscan(in, &kval)

			allUnit[i] = kval
		}
		var min_dev int
		sortUnit = append(sortUnit, 0)
		min_dev = int(math.Abs(float64(allUnit[0]) - float64(allUnit[1])))
		for i = 1; i < nUnit-1; i++ {
			//	var min_dev int
			if allUnit[i] > 0 {
				if i%2 == 0 {
					sortUnit = append(sortUnit, i)
				}

				min_dev = int(math.Abs(float64(allUnit[i]) - float64(allUnit[i+1])))
			}
			fmt.Println(min_dev)
			for j = i + 1; j < nUnit; j++ {

				if allUnit[j] > 0 && (min_dev > allUnit[i]-allUnit[j]) {
					min_dev = int(math.Abs(float64(allUnit[i]) - float64(allUnit[j])))
					sortUnit = append(sortUnit, j)
					allUnit[j] = 0
				}
			}
			allUnit[i] = 0
		}

		fmt.Println(allUnit)
	}
	fmt.Println(sortUnit)

}
