package main

import "fmt"

func main() {
	rating := []int{2, 5, 3, 4, 1}
	res := numTeams(rating)
	fmt.Println("res: ", res)
}

func numTeams(rating []int) int {
	ans := 0
	for i := 0; i < len(rating); i++ {
		for j := i; j < len(rating); j++ {
			for k := j; k < len(rating); k++ {
				if rating[i] > rating[j] && rating[j] > rating[k] {
					ans++
				} else if rating[i] < rating[j] && rating[j] < rating[k] {
					ans++
				}
			}
		}
	}
	return ans
}
