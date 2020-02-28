/*
这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。

请你判断自己是否能够跳到对应元素值为 0 的 任意 下标处。

注意，不管是什么情况下，你都无法跳到数组之外。



示例 1：

输入：arr = [4,2,3,0,3,1,2], start = 5
输出：true
解释：
到达值为 0 的下标 3 有以下可能方案：
下标 5 -> 下标 4 -> 下标 1 -> 下标 3
下标 5 -> 下标 6 -> 下标 4 -> 下标 1 -> 下标 3
示例 2：

输入：arr = [4,2,3,0,3,1,2], start = 0
输出：true
解释：
到达值为 0 的下标 3 有以下可能方案：
下标 0 -> 下标 4 -> 下标 1 -> 下标 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "container/list"

//深搜
// 边界值 1 出圈，2，成环或扫过，3，为0

func canReach(arr []int, start int) bool {
	markedList := make([]int, len(arr))
	return dfs(arr, start, markedList)
}

func dfs(arr []int, pos int, markedList []int) bool {
	if pos < 0 || pos >= len(arr) {
		return false
	}
	if markedList[pos] == 1 || markedList[pos] == 2 {
		return false
	}
	markedList[pos] = 1
	if arr[pos] == 0 {
		return true
	}
	if dfs(arr, pos+arr[pos], markedList) || dfs(arr, pos-arr[pos], markedList) {
		return true
	}
	return false
}

//广搜
// 边界值 1 出圈，2，扫过，3，

func canReach(arr []int, start int) bool {
	marketList := make([]int, len(arr))
	queue := list.New()
	queue.PushBack(start)
	marketList[start] = 1
	for queue.Len() > 0 {
		tmp := queue.Front()
		num := tmp.Value.(int)
		queue.Remove(tmp)
		if arr[num] == 0 {
			return true
		}
		if num+arr[num] < len(arr) && marketList[num+arr[num]] == 0 {
			marketList[num+arr[num]] = 1
			queue.PushBack(num + arr[num])
		}
		if num-arr[num] >= 0 && marketList[num-arr[num]] == 0 {
			marketList[num-arr[num]] = 1
			queue.PushBack(num - arr[num])
		}
	}
	return false
}
