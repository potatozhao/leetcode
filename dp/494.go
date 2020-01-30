/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释: 

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
注意:

数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/


package main


// 记忆化递归
func findTargetSumWays(nums []int, S int) int {
    dpMap := make(map[string]int)
    return fundtage(nums, 0, S, dpMap)
}

func fundtage(nums []int, start int, num int, dpmap map[string]int) int{
    if start == len(nums) && num == 0{
        return 1
    }else if start == len(nums){
        return 0
    }
    key := fmt.Sprintf("%d_%d", start, num)
    if ret, ok := dpmap[key]; ok{
        return ret
    }
    big := fundtage(nums, start + 1, num + nums[start], dpmap)
    small := fundtage(nums, start + 1, num - nums[start], dpmap)
    dpmap[key] = big + small
	return big + small
}


/*

该题难点在于，如何将其转换为背包问题。

假设有集合 N 和 集合 P 使得 sum(N) - sum(P) = S 其中 N，P 均是输入数组的子集, 那么必然有

2*sum(N) - sum(P) + sum(P) = S + sum(N) + sum(P)
=>
2*sum(N) = S + sum(nums)
所以我们只需要找到满足上述推论的集合 N 就 ok 了，即转换为问题：存在一个容量为 V 的背包，从 nums 中任意抽取一定数量的数，使得背包恰好被放满，有多少种放法。这个容量 V 我们通过上述的推导已经求出来了, V = sum(N)。

作者：nfgc
链接：https://leetcode-cn.com/problems/target-sum/solution/zhuan-huan-wei-bei-bao-wen-ti-yi-wei-er-wei-liang-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

//01背包解法
func findTargetSumWays_01(nums []int, S int) int {
    V := 0
    for _,i := range nums{
        V += i
    }
    if V < S || (S+V) % 2 ==1{
        return 0
    }
    V += S
    V /= 2
    dpList := make([]int, V+1)
    dpList[0] = 1
    for _, v := range nums{
        for i := V; i >=v; i--{
                dpList[i] += dpList[i - v]
        }
    }
    return dpList[V]
}