package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const MAXTASK = 50000

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nTest, i, nTask, taskNum, n int

	fmt.Fscan(in, &nTest)

	for n = 0; n < nTest; n++ {

		fmt.Fscan(in, &nTask)
		flag := true
		task := make([]int, nTask, nTask)
		taskMax := make([]int, MAXTASK, MAXTASK)

		for i = 0; i < nTask; i++ {
			fmt.Fscan(in, &task[i])
		}

		for i = 0; i < nTask; i++ {
			taskNum = task[i]

			if taskMax[taskNum] == 1 && i > 0 && task[i-1] != taskNum {
				flag = false
			}
			taskMax[taskNum] = 1
		}

		if flag {
			fmt.Println("YES")
		} else if !flag {
			fmt.Println("NO")
		}
	}
}
