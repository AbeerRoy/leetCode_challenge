package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{0, 0, 1, 3, 4, 0, 4, 5}
	//nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))

}

// only remove values/elements that repeated more than twice
func removeDuplicates(nums []int) []int {
	sort.Ints(nums)

	//slices.Sort(nums)
	flags := map[int]bool{}
	// counters 1
	i := 0
	//memory address
	l := 0

	for value := range nums {
		i = value
		if !flags[nums[value]] {
			flags[nums[value]] = true
			if i < len(nums)-1 {
				//counter 2
				j := i + 1
				//fmt.Println(j)
				if nums[i] == nums[j] {
					//fmt.Println(nums[j])
					//for minimum reapeating value of 2
					for k := 0; k < 2; k++ {
						nums[l] = nums[value]
						//fmt.Println(nums[l])
						l++
					}
				} else {
					nums[l] = nums[value]
					//fmt.Println(nums[l])
					l++
				}
			}
			if i == len(nums)-1 {
				nums[l] = nums[value]
				//fmt.Println(nums[l])
				l++
			}
		}
	}
	nums = nums[:l]
	return nums
}
