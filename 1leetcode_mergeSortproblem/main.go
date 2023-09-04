package main

import (
	"fmt"
	"sort"
)

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	const m int = 3
	const n int = 3

	fmt.Println("\n--- Sorted ---\n\n", merge(nums1, m, nums2, n), "\n")
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 1 && len(nums1) == 1 {
		copy(nums1, nums2)
	}

	if m == 1 && n == 1 {
		copy(nums1[1:], nums2)
	}

	if m == 1 && n != 1 {
		copy(nums1[1:], nums2)
	}

	if m < 1 {
		copy(nums1, nums2)
	}
	if m > 1 {
		copy(nums1[m:], nums2)
	}

	sort.Ints(nums1)

	return nums1
}
