/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
示例 2:

输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// 股票有三种状态，1.波动，2.连续增，3连续跌，1在波低买入，波峰卖出，2 买，3不动
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	num := 0
	a := -1
	b := -1
	for i := 0; i < len(prices); i++ {
		//取波底作为a
		if (prices[i] < a || a < 0) && b < 0 {
			a = prices[i]
		} else if prices[i] > b && a < prices[i] && a >= 0 { //取波峰作为b
			b = prices[i]
		} else if a >= 0 && b >= 0 && prices[i] < b { //跌了就卖
			num += b - a
			a = prices[i]
			b = -1
		}
		if i == len(prices)-1 && a >= 0 && b >= 0 { //连续涨的话，结尾强卖
			num += b - a
		}
	}
	return num
}

// 贪心算法，只加正数！我勒个去。。。这啥算法。。。
func maxProfit_2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	num := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			num += prices[i+1] - prices[i]
		}
	}
	return num
}

// dp解法 每一个股票有两个状态，dp[i][0-1], i为钱在手，1为换成股票了。。。 值为收益。
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	a := 0
	b := 0 - prices[0]
	for i := 1; i < len(prices); i++ {
		a = max(a, b+prices[i]) // 前为当前不买不卖，后为把股票换成钱
		b = max(b, a-prices[i]) // 前为当前不买不卖，后为把钱换成股票
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
