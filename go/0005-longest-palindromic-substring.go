package main

import "fmt"

func checkPalindromic(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		// fmt.Println(string(s[left]), string(s[right]))
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func longestPalindrome(s string) string {
	for size := len(s); size > 0; size-- {
		for start := 0; start <= len(s)-size; start++ {
			sub_str := s[start : start+size]
			// fmt.Println("sub str", sub_str, "check", checkPalindromic(sub_str))
			if checkPalindromic(sub_str) {
				return sub_str
			}
		}
	}
	return ""
}

func make2DArray(row int, col int) [][]bool {
	arr := make([][]bool, row)
	for i := range arr {
		arr[i] = make([]bool, col)
	}
	return arr
}

func showMatrix(matrix [][]bool) {
	fmt.Println("DP array:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] {
				fmt.Print(" 0 ")
			} else {
				fmt.Print(" X ")
			}
		}
		fmt.Println()

		// fmt.Println(matrix[i])
	}

}

func longestPalindromeDynamicProg(s string) string {
	n := len(s)
	dp := make2DArray(n, n)

	for i := range n {
		dp[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
	}

	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			fmt.Printf("[%d][%d]: %s, %s\n", i, j, string(s[i]), string(s[j]))
			last_val := dp[i+1][j-1]
			is_this := s[i] == s[j]
			dp[i][j] = last_val && is_this
		}
	}

	// Print the dp array
	showMatrix(dp)

	return s

}

func testLongestPalindrome() {
	// s := "babad"
	s := "cbabca"
	// s := "cbbb"
	// s := "azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc"
	// s := "a"
	// s := "cbabaa"

	// res := longestPalindrome(s)
	res := longestPalindromeDynamicProg(s)
	fmt.Println("res:", res)
}
