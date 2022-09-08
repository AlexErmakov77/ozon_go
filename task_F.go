package main

import "fmt"

func user_in_set(user int, set map[int]bool) bool {
	_, ok := set[user]
	return ok
}

func count_common_friends(friends1 map[int]bool, friends2 map[int]bool) int {
	count := 0
	for f := range friends1 {
		if user_in_set(f, friends2) {
			count += 1
		}
	}
	return count
}

func main() {
	var users, pairs int
	fmt.Scanf("%d%d", &users, &pairs)

	var friends = map[int]map[int]bool{}
	for i := 1; i <= users; i++ {
		friends[i] = map[int]bool{}
	}

	for i := 0; i < pairs; i++ {
		var u1, u2 int
		fmt.Scanf("%d%d", &u1, &u2)
		friends[u1][u2] = true
		friends[u2][u1] = true
	}

	fmt.Println()
	fmt.Println("Friends for each user:")
	for u := 1; u <= users; u++ {
		fmt.Print(u, ": [ ")
		for friend, value := range friends[u] {
			if value {
				fmt.Print(friend, " ")
			}
		}
		fmt.Println("]")
	}

	fmt.Println()
	fmt.Println("Answer:")
	for ux := 1; ux <= users; ux++ {
		fmt.Print(ux, ": ")
		res := []int{}
		maxCount := 0
		for uy := 1; uy <= users; uy++ {
			if ux != uy && !user_in_set(uy, friends[ux]) {
				var count = count_common_friends(friends[ux], friends[uy])
				//        fmt.Print(uy, "(", count, ") ")
				if count > maxCount {
					maxCount = count
					res = nil
					res = append(res, uy)
				} else if count > 0 && count == maxCount {
					res = append(res, uy)
				}
			}
		}
		//    fmt.Println()
		res_len := len(res)
		for i := 0; i < res_len; i++ {
			fmt.Print(res[i], " ")
		}
		if res_len == 0 {
			fmt.Print(0)
		}
		fmt.Println()
	}
}
