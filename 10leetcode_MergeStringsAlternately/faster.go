package main

import (
	"fmt"
)

func main() {
	word1 := "abcde"
	word2 := "pqrs"
	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	var left = ""
	// Word1 length
	var i1 = 0
	// Word2 length
	var i2 = 0

	for i := 0; i < (len(word1) + len(word2)); i++ {
		// word2 length equal to word 1
		if len(word1) == len(word2) {
			if i%2 == 0 {
				left = left + string(word1[i1])
				i1++
			} else if i%2 == 1 {
				left = left + string(word2[i2])
				i2++

			}
			// word1 length larger than word2
		} else {
			if len(word1) > len(word2) {
				if len(word2)*2 <= i {
					left = left + string(word1[i1])
					i1++
				} else {
					if i%2 == 0 {
						left = left + string(word1[i1])
						i1++
					} else if i%2 == 1 {
						left = left + string(word2[i2])
						i2++
					}

				}
				// word2 length larger than word 1
			} else {
				if len(word1)*2 <= i {
					left = left + string(word2[i2])
					i2++

				} else {
					if i%2 == 0 {
						left = left + string(word1[i1])
						i1++
					} else if i%2 == 1 {
						left = left + string(word2[i2])
						i2++
					}

				}

			}

		}
	}

	return left

}
