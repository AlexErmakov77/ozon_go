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

	var nInput, nTovarov, i, j int32
	var summa int32

	fmt.Fscan(in, &nInput)
	result := make([]int32, nInput, nInput)

	for j = 0; j < nInput; j++ {

		fmt.Fscan(in, &nTovarov)
		pokupki := make([]int32, nTovarov, nTovarov)

		mapKorzina := make(map[int32]int32)

		for i = 0; i < nTovarov; i++ {
			var tovar int32
			fmt.Fscan(in, &tovar)
			pokupki[i] = tovar
		}

		for _, v := range pokupki {
			mapKorzina[v]++
		}
		summa = 0

		for k, v := range mapKorzina {

			if v/3 > 0 {
				summa = summa + (v-v/3)*k
			} else {
				summa = summa + k*v
			}

		}
		//	fmt.Println(summa)
		result[j] = summa
	}

	for j = 0; j < nInput; j++ {
		fmt.Println(result[j])
	}
}
