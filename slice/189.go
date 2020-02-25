/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func rotate(nums []int, k int) {
	solve2(nums, k)
}

func solve1(nums []int, k int) {
	tmpNums := make([]int, len(nums))
	for i := range nums {
		tmpNums[(i+k)%len(nums)] = nums[i]
	}
	for i := range nums {
		nums[i] = tmpNums[i]
	}
}

// 三次翻转法
func solve2(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	overturn(nums[:len(nums)-k])
	overturn(nums[len(nums)-k:])
	overturn(nums)
}

func overturn(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

// 证明共识
/*
题目要求对于一个给定长度数组完成下标的映射 这里讨论移动前后的下标对应关系
满足区间函数 ，这里简化k为 0＜k＜n,实际超过的部分可以通过取模运算映射到这个范围

f(x)=x+k ,（0≤x≤n-1-k）
f(x)=x+k-n,(x＞n-1-k)
对于第一次翻转
0≤x≤n-1-k,f1(x)满足

f1(x)+x = n-1-k
化简得到 f1(x)=n-1-k-x 此时f1(x)≥0
第二次翻转(x＞n-1-k)

f2(x)+x=n-1+n-k
化简得到f2(x)=2n-1-k-x
第三次翻转
0≤x≤n-1-k时

f(x)+f1(x)=n-1
代入化简 f(x)+n-1-k-x = n-1 即f(x)=x+k;
x＞n-1-k时

f(x)+f2(x)=n-1
代入化简 f(x)+2n-1-k-x = n-1 即f(x)=x+k-n;
正好符合右移的映射关系
说明三次翻转操作后所得数组即为右移后的结果

作者：夕枫晚照
链接：https://www.jianshu.com/p/c6b19d5847ad
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
