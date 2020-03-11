/*
 */

package main

// 内存炸了。
func canJump(nums []int) bool {
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		tmpList := make([]int, 0)
		for i := range queue {
			if queue[i] == len(nums)-1 {
				return true
			}
			for j := queue[i] + 1; j <= queue[i]+nums[queue[i]]; j++ {
				if j == len(nums)-1 {
					return true
				}
				if j < len(nums) && nums[j] != 0 {
					tmpList = append(tmpList, j)
				}
			}
		}
		queue = tmpList
	}
	return false
}

//dfs 差点超时
func canJump(nums []int) bool {
	markedList := make([]bool, len(nums))
	return dfs(nums, 0, markedList)
}

func dfs(nums []int, pos int, markedList []bool) bool {
	if markedList[pos] {
		return false
	}
	if pos == len(nums)-1 {
		return true
	}
	if nums[pos] == 0 {
		return false
	}
	for i := pos + 1; i < len(nums) && i <= pos+nums[pos]; i++ {
		ret := dfs(nums, i, markedList)
		if ret {
			return true
		}
	}
	markedList[pos] = true
	return false
}

//神他妈解法！
/*
解题思路：

如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点。
可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新。
如果可以一直跳到最后，就成功了。

作者：ikaruga
链接：https://leetcode-cn.com/problems/jump-game/solution/55-by-ikaruga/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func canJump(nums []int) bool {
	maxNum := 0
	for i := 0; i <= maxNum; i++ {
		if maxNum >= len(nums)-1 {
			return true
		}
		maxNum = max(maxNum, i+nums[i])
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
