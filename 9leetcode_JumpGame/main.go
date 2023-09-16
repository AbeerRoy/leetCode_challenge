package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4} //true
	//nums := []int{2, 5, 0, 0} //false
	//nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))

}
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	} else if nums[0] > 0 {
		var i = nums[0]
		end := &nums[len(nums)-1]

		for {
			if i >= len(nums) || i == len(nums)-1 {
				return true

			} else if &nums[i] == end {
				return true

			} else if nums[i] == 0 {
				return false
			} else {
				if i+nums[i] < len(nums) {
					i = i + nums[i]

				} else {
					return true
				}
			}

		}
	} else {
		return false
	}

}
