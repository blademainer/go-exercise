package main

import (
	"fmt"
)

func main() {
	groupSize := 0
	_, err := fmt.Scanf("%d", &groupSize)
	if err != nil {
		panic(err)
	}
	for i := 0; i < groupSize; i++ {
		var n, m int
		_, err := fmt.Scanf("%d %d", &n, &m)
		if err != nil {
			panic(err)
		}
		comments := make([][]int, n)
		for j := 0; j < n; j++ {
			var a, b int
			_, err := fmt.Scanf("%d %d", &a, &b)
			if err != nil {
				panic(err)
			}
			comments[j] = []int{a, b}
		}

		friends := make(map[int]struct{})
		for j := 0; j < m; j++ {
			friend := 0
			_, err = fmt.Scanf("%d", &friend)
			if err != nil {
				panic(err)
			}
			friends[friend] = struct{}{}
		}
		fmt.Print(minFriends(comments, friends))
		if i != groupSize-1 {
			fmt.Println()
		}
	}
}

func minFriends(comments [][]int, friends map[int]struct{}) int {
	stat := make(map[int]map[int]int) // a -> b -> count
	for _, comment := range comments {
		a, b := comment[0], comment[1]

		_, aok := friends[a]
		_, bok := friends[b]
		// 如果 a 和 b 都不是好友，那么这条评论对小 Q 是不可见的
		if !aok && !bok {
			continue
		} else if aok && bok {
			continue
		}

		if aok {
			if statb, ok := stat[a]; ok {
				statb[b]++
			} else {
				stat[a] = map[int]int{b: 1}
			}
		}

		if bok {
			if statb, ok := stat[b]; ok {
				statb[a]++
			} else {
				stat[b] = map[int]int{a: 1}
			}
		}

	}

	max := 0
	// goodFriend := friends[0]
	for _, m2 := range stat {
		for _, i3 := range m2 {
			if i3 > max {
				max = i3
				// goodFriend = i2
			}
		}
	}
	return max + len(comments)
}
