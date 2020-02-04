/*
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
示例:

输入: [1,2,3,0,2]
输出: 3 
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
*/

package main

// dp解法
func maxProfit(prices []int) int {
    if len(prices) <=1{
        return 0
    }
    dpList := make([]int, len(prices) + 2)
    for j := 3; j< len(prices) + 2; j++{
        for i := j - 1; i >=2; i--{
                dpList[j] = max(dpList[j], dpList[i], dpList[i-2] + prices[j-2] - prices[i-2])
        }
    }
    return dpList[len(prices) + 1]
}

func max(a, b, c int) int{
    if a >= b && a >= c{
        return a
    }else if b >= a && b >= c{
        return b
    }
    return c
}

// 套路解法
func maxProfit_2(prices []int) int {
    const intMax = int(^uint(0) >> 1) // int最大值
    const intMin = ^intMax // int最小值

    a := 0 // 不持有的资产
    b := intMin // 持有的资产
    c := 0 // 解冻前不持有的资产 解冻期为1 所以，c := a - 1

    for i := 0; i<len(prices); i++{
        tmp := a
        a = max(a, b + prices[i])
        b = max(b, c - prices[i])
        c = tmp
    }
    return a
}

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}