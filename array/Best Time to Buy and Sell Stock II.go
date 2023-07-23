package array

func maxProfitII(prices []int) int {
	profit := 0

	//便利报价：只要第i天的价格比i-1天高，那我们将价格差记录到总利润中
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
