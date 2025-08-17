package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left := 0
	max_length := 0
	m := make(map[rune]int)
	// for right := 0; right < len(s); right++ {
	for right, r := range s {
		fmt.Printf("[%d]: %c\n", right, r)
		m[r]++
		for {
			if m[r] > 1 {
				left++
				m[rune(s[left])]--
			} else {
				break
			}
		}

		max_length = max(max_length, right-left+1)
	}
	return max_length
}

func isCharInList(list []rune, char rune) int {
	for i, c := range list {
		if char == c {
			return i
		}
	}
	return -1

}

// find with sliding window
func lengthOfLongestSubstring2(s string) int {
	left := 0
	max_length := 0
	list := []rune{}
	for right, char := range s {
		index_in_list := isCharInList(list, char)
		if index_in_list == -1 {
			list = append(list, char)
			fmt.Printf("the %d ,%c added to the list \n", right, char)
		} else {
			// the character is in the list
			// 1. we should cut the list to the character index (including itself)
			// 2. update left to be left + index_in_list (offset)
			fmt.Printf("the %d, %c is already in the list, index: %d\n", right, char, index_in_list)
			list = list[index_in_list+1:]
			list = append(list, char)
			left = left + index_in_list + 1
		}
		diff := right - left + 1
		if max_length < diff {
			max_length = diff
		}

	}
	return max_length

}
