package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums := []int{1, 2, 3, 5}
	k := 6
	fmt.Println(rotate(nums, k))

}

func rotate(nums []int, k int) []int {

	length := len(nums)
	l := 0

	// when rotational value is
	if k > length {
		counter := 1
		rotated := make([]int, k)
		for i := k; i > k-length; i-- {
			rotated[l] = nums[length-counter]
			//fmt.Println(i)
			l++
			counter++
		}
		for i := 0; i < len(rotated); i++ {
			if rotated[i] > 0 {
				nums[i] = rotated[i]
				//fmt.Println(nums)

			}

			// }
			// rotated := make([]int, length+1)
			// for i := 1; i < len(rotated)-1; i++ {
			// 	rotated[counter-i] = nums[length-i]
			// 	counter = counter + 2
			// 	length--
			// }
		}

		return rotated

	} else {
		// when there is only one element in the slice
		if length <= 1 {
			return nums
		} else {
			//when slice length and rotational value is equal
			if k == length {
				counter := 1
				rotated := make([]int, k)
				for i := k; i > k-length; i-- {
					rotated[l] = nums[length-counter]
					//fmt.Println(i)
					l++
					counter++
				}
				for i := 0; i < len(rotated); i++ {
					if rotated[i] > 0 {
						nums[i] = rotated[i]
						//fmt.Println(nums)

					}

				}

				return nums

			} else {
				// rotational value less than slice length

				rotated := make([]int, length)
				for i := k; i > 0; i-- {
					rotated[l] = nums[length-i]
					fmt.Println(i)
					l++
				}
				for i := 0; i < length-k; i++ {
					rotated[l] = nums[i]
					//fmt.Println(i)
					l++
				}
				for i := 0; i < len(rotated); i++ {
					nums[i] = rotated[i]
					//fmt.Println(nums)
				}
				return nums

			}

		}

	}
}
