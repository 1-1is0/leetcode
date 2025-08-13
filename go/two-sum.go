package main

// question link: https://leetcode.com/problems/two-sum/description/
import "fmt"

func twoSum(nums []int, target int) []int {
	for i, n1 := range nums {
		for j, n2 := range nums[i+1:] {
			// fmt.Println("i:", i, "j:", j, "n1:", n1, "n2:", n2)
			if n1+n2 == target {
				return []int{i, i + j + 1}
			}
		}
		// fmt.Println("next i")
	}
	return []int{}
}

func twoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n1 := range nums {
		m[n1] = i
	}
	for i, n1 := range nums {
		other_half := target - n1
		if j, ok := m[other_half]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9
	nums := []int{3, 2, 4}
	target := 6
	// fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumHash(nums, target))
}
