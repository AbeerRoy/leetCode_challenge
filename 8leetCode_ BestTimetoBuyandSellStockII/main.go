package main

import "fmt"

func main() {
	prices := []int{}

	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	var profit = 0
	//var profit []int
	//var minimumPrice = prices[0]
	var j = 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[j] {
			value := prices[i] - prices[j]
			profit += value

			//profit = append(profit, value)

		}
		j++
	}
	return profit

}
