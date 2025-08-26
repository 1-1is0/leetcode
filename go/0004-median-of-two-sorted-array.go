package main

import (
	"fmt"
)

func findMedian(nums []int) (int, float64) {
	if len(nums) == 0 {
		return -1, 0.0
	}
	medianIndex := int(len(nums) / 2)
	median := 0.0
	if len(nums)%2 == 0 {
		// fmt.Printf("nums len %d, nums: %s\n", len(nums), fmt.Sprint(nums))
		median = (float64(nums[medianIndex-1]) + float64(nums[medianIndex])) / 2.0
	} else {
		median = float64(nums[medianIndex])
	}
	return medianIndex, median

}

func getRightHalf(list []int) []int {
	midIndex := int(len(list) / 2)
	if len(list)%2 == 0 {
		return list[midIndex:]
	} else if len(list) <= 1 {
		return list
	}
	return list[midIndex+1:]
}

func getLeftHalf(list []int) []int {
	midIndex := int(len(list) / 2)
	if len(list)%2 == 0 {
		return list[:midIndex]
	} else if len(list) <= 1 {
		return list
	}
	return list[:midIndex+1]

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Println("##########        findMedianSortedArrays           #########")
	medIndex1, med1 := findMedian(nums1)
	medIndex2, med2 := findMedian(nums2)
	if medIndex1 == -1 {
		fmt.Println("return med2, 1 is Empty")
		return med2
	} else if medIndex2 == -1 {
		fmt.Println("return med1, 2 is Empty")
		return med1
	}
	if med1 == med2 {
		fmt.Printf("return med1: %f, med2: %f\n", med1, med2)
		return med1
	} else if len(nums1) <= 1 && len(nums2) <= 1 {
		fmt.Printf("return med1: %f, med2: %f\n", med1, med2)
		return (med1 + med2) / 2
	} else if med1 < med2 {
		nums1 = getRightHalf(nums1)
		nums2 = getLeftHalf(nums2)
		fmt.Printf("num1 < num2 --- nums1 %s, nums2: %s\n", fmt.Sprint(nums1), fmt.Sprint(nums2))
		return findMedianSortedArrays(nums1, nums2)
	} else {
		nums1 = getLeftHalf(nums1)
		nums2 = getRightHalf(nums2)
		fmt.Printf("num1 > num2 ---- nums1 %s, nums2: %s\n", fmt.Sprint(nums1), fmt.Sprint(nums2))
		return findMedianSortedArrays(nums1, nums2)
	}
}

func testFindMedianSortedArrays() {
	// nums1 := [1,3]
	// nums2 := [2]

	// nums1 := []int{1, 2}
	// nums2 := []int{3, 4}

	// nums1 := []int{1, 3}
	// nums2 := []int{2}

	// nums1 := []int{}
	// nums2 := []int{1}

	nums1 := []int{2, 2, 4, 4}
	nums2 := []int{2, 2, 2, 4, 4}

	// res1 := getRightHalf(nums2)
	// fmt.Println("res1", res1)
	// res2 := getRightHalf(res1)
	// fmt.Println("res2", res2)
	// res3 := getRightHalf(res2)
	// fmt.Println("res3", res3)

	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println("res", res)
}

// 1: 3  2: 2  1 > 2
// new nums1: [4, 4]        nums2: [2, 2]
// 1: 4       2: 2      1 > 2
// new nums1: [4]          nums2: [2]
// 1: 3.5     2: 2      1 > 2
// news nums1: [4]     nums2: [2]         -> (2 + 4) /2 = 3

// 2, 2, 2, 2, 2, 4, 4, 4

// nums1: [2, 2, 4, 4]        nums2: [2, 2, 2, 4, 4]     m1: 3.5     m2: 2     m1 > m2     or m2 < m1

// nums1: [2, '2, 4']         nums2: ['2', 4, 4]         m1: 2    m2: 4        m1 < m2

// nums1: ['2', 4]           nums2: [2, '4']             m1: 3   m2: 3 return 3

// nums1 []

// nums1: [2, 2, 4, 4]        nums2: [2, 2, 2, 4, 4]     m1: 3     m2: 2     m1 > m2     or m2 < m1

// nums1: [2, 2]         nums2: [4, 4]         m1: 2    m2: 4        m1 < m2

// nums1: [2]           nums2: [4]             m1: 3   m2: 3 return 3

// nums1 []
