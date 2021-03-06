/*
现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。

你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。

注意:

给定 0 和 1 的数量都不会超过 100。
给定字符串数组的长度不会超过 600。
示例 1:

输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
输出: 4

解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。
示例 2:

输入: Array = {"10", "0", "1"}, m = 1, n = 1
输出: 2

解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ones-and-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


package main


// 多维01背包问题，
func findMaxForm(strs []string, m int, n int) int {
	zMap := make(map[string]int)
	iMap := make(map[string]int)
	for _, v := range strs {
		zNum := 0
		iNum := 0
		for i, _ := range v {
			if v[i] == '0' {
				zNum += 1
			} else {
				iNum += 1
			}
		}
		zMap[v] = zNum
		iMap[v] = iNum
	}
	MNList := make([][]int, m+1)
	for m, _ := range MNList {
		MNList[m] = make([]int, n+1)
	}
	for _, v := range strs {
		vm := zMap[v]
		vn := iMap[v]
		for i := m; i >= vm; i-- {
			for j := n; j >= vn; j-- {
				if i >= zMap[v] && j >= iMap[v] {
					MNList[i][j] = max(MNList[i][j], MNList[i-vm][j-vn]+1)
				}
			}
		}
	}
	return MNList[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
