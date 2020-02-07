/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func maxProfit(k int, prices []int) int {
    if k <=0 || len(prices) <=1{
        return 0
    }
    intMax := int(^uint(0) >> 1)
    intMin := ^intMax
    if k >> 1 > len(prices){ // 如果K > 1/2len(prices)，也就意味着k是用不完的，也就意味着K是无限大的。
        a:=0
        b:=intMin
        for i :=0; i<len(prices); i++{
            tmp := a
            a = max(a, b + prices[i])
            b = max(b, tmp - prices[i])
        }
        return a
    }else{
        dpList := make([][2]int, k+1)
        for i:= 1; i<= k; i++{
            dpList[i][1] = intMin
        }
        for i := 0; i<len(prices);i++{
            for j :=1; j<=k; j++{
                dpList[j][0] = max(dpList[j][0], dpList[j][1] +prices[i])
                dpList[j][1] = max(dpList[j][1], dpList[j-1][0] - prices[i])
            }
        }
        return dpList[k][0]
    }
}

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}