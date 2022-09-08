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

	sortUnit := []int{}

	for j = 0; j < nInput; j++ {

		fmt.Fscan(in, &nUnit)

		allUnit := make([]int, nUnit, nUnit)

		for i = 0; i < nUnit; i++ {
			var kval int
			fmt.Fscan(in, &kval)

			allUnit[i] = kval
		}
		var min_deviation, count int

		for j = 0; j < nUnit-1; j++ {

			if allUnit[j] > 0 {
				min_deviation = int(math.Abs(float64(allUnit[j]) - float64(allUnit[j+1])))

				sortUnit = append(sortUnit, j)

				for i = j + 1; i < nUnit; i++ {
					if allUnit[i] > 0 && min_deviation > int(math.Abs(float64(allUnit[j])-float64(allUnit[i]))) {
						min_deviation = int(math.Abs(float64(allUnit[j]) - float64(allUnit[i])))
						count = i
						sortUnit = append(sortUnit, count)
					}

					allUnit[count] = 0
				}
				allUnit[j] = 0
			}
		}
		fmt.Println(allUnit)
	}
	fmt.Println(sortUnit)

}
