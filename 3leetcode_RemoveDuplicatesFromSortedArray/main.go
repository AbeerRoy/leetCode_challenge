package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 3, 8, 9}
	fmt.Println(removeDuplicates(nums))

}

func removeDuplicates(nums []int) []int {

	

	// l := 1
	// for r := 1; r < len(nums); r++ {
	// 	if nums[r] != nums[r-1] {
	// 		nums[l] = nums[r]
	// 		l++
	// 	}
	// }
	// return nums[:l]
	/* 
	
			**** Very high Complexity **** 
	
	*/
	//l := 0
	// var length = len(nums)
	// for i := 0; i < length; i++ {
	// 	for j := i + 1; j < length; j++ {
	// 		if nums[i] == nums[j] {
	// 			l++
	// 			for k := j; k < length-1; k++ {
	// 				fmt.Println(nums[k+1])
	// 				nums[k] = nums[k+1]

	// 			}
	// 			length--
	// 			j--
	// 		}

	// 	}
	// }
	/* 
			**** Using Map to store flag,memory intensive ****
	*/

	flags := map[int]bool{}
	i := 0
	for j := range nums {
		if !flags[nums[j]] {
			flags[nums[j]] = true
			nums[i] = nums[j]
			i++

		}
	}
	nums = nums[:i]
	return nums

}
