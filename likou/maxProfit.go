package likou

import "math"

//121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	mixPrice := math.MaxInt64
	maxPrice := 0

	for i := range prices {
		if prices[i] < mixPrice {
			mixPrice = prices[i]
		} else if prices[i]-mixPrice > maxPrice {
			//最大价格
			maxPrice = prices[i] - mixPrice
		}
	}
	return maxPrice
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	mixPrice := math.MaxInt64
	maxPrice := 0

	for i := range prices {
		maxPrice = max(maxPrice, prices[i]-mixPrice)
		mixPrice = mix(mixPrice, prices[i])
	}
	return maxPrice
}
