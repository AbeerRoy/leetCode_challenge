package main

import (
	"fmt"
)

func main() {
	//haystack := "sadbutsad"
	//needle := "sadsad"
	haystack := "mississippi"
	needle := "issip"
	//needle := "pi"
	//haystack := "babba"
	//needle := "bb"
	//haystack := "aabaaabaaac"
	//needle := "aabaaac"
	//haystack := "aaaabbbababb"
	//needle := "baa"

	fmt.Println(Hold(haystack, needle))

}
func Hold(haystack string, needle string) int {
	//stores length of needle
	l1 := len(needle)
	//sets an index counter for needle
	l := 0

	for i := 0; i < len(haystack); i++ {
		//check if each character matches
		if haystack[i] == needle[l] {

			//Best case: checks if all the charater matched
			if l == l1-1 {
				return i - l
			}
			//increaments index
			l++

		} else {
			//if haystack[i] != needle[l]
			//copies the value of index of haystack to find the next value of needle[l] which didn't match in Best Case
			var k = i
			for k = i; k < len(haystack); k++ {

				/***********  Here first value of needle and the last value of the needle is crosschecked with haystack string before exiting the nested loop   ***********/
				// setting all values inbetween as bool absolute could ensure ULTIMATE final Check before exiting the loop
				c := haystack[k-l : k+1]
				fmt.Println(c)

				// setting all values of needle as bool absolute to ensure ULTIMATE final Check before exiting the loop Second Best Case
				if haystack[k-l:k+1] == needle {
					return k - l

				}
				if haystack[k] == needle[l] && haystack[k-l] == needle[0] {
					//reseting the value of i(original index of haystack)
					i = k - l - 1
					//reseting the value of l for first loop
					l = 0
					//exit from nested loop
					break

				}

			}
			//if no match found, Best case for exit
			if k == len(haystack) {
				return -1

			}

		}

	}
	//if no match found,second base case for exit
	return -1

}
