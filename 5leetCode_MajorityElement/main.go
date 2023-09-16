package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(majorityElement(nums))

}

func majorityElement(nums []int) int {
	var sortedList = nums
	sort.Ints(sortedList)
	var e = len(nums) / 2
	return nums[e]
}
