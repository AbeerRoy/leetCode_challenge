package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	length := len(prices)
	//minimum price
	var price = prices[0]
	//temporary profit variable
	var profit = 0

	//profitCheck

	if length == 0 {
		return 0
	} else if length == 1 {
		return profit

	} else {
		for i := 1; i < length; i++ {
			if prices[i] < price {
				price = prices[i]
			} else if (prices[i] - price) > profit {
				profit = prices[i] - price

			}

		}

		return profit
	}

}
