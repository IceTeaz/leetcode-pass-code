package main

func main() {

}

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	return dp[amount] == amount+1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
