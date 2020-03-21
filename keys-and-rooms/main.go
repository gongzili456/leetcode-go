// https://leetcode-cn.com/problems/keys-and-rooms/
// 841. 钥匙和房间
package main

import "fmt"

func main() {
	rooms := [][]int{
		// {1}, {2}, {3}, {},
		{1, 3}, {3, 0, 1}, {2}, {0},
	}
	res := canVisitAllRooms(rooms)
	fmt.Println("res: ", res)
}

func canVisitAllRooms(rooms [][]int) bool {
	Q := [][]int{}
	V := map[int]bool{0: true}

	Q = append(Q, rooms[0])

	for len(Q) > 0 {
		rs := Q[0]
		Q = Q[1:]
		for _, r := range rs {
			if _, ok := V[r]; ok {
				continue
			}
			V[r] = true

			Q = append(Q, rooms[r])
		}
	}

	return len(V) == len(rooms)
}
