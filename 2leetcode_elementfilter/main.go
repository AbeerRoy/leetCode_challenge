package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 4}
	val := 4
	fmt.Println(removeElement(nums, val))

}
func removeElement(nums []int, val int) int {
	i := 0
	// for j := 0; j < len(nums); j++ {
	// 	if nums[j] != val {
	// 		nums[i] = nums[j]
	// 		i++
	// 	}
	for j, value := range nums {
		if value != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
