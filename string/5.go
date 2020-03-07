/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	max := ""
	for i := 1; i < len(s); i++ {
		x, y := i, i
		temp := check(s, x, y)
		if len(max) < len(temp) {
			max = temp
		}
		x, y = i-1, i
		temp = check(s, x, y)
		if len(max) < len(temp) {
			max = temp
		}
	}
	return max
}

func check(s string, x, y int) string {
	for x >= 0 && y < len(s) && s[x] == s[y] {
		x--
		y++
	}
	return s[x+1 : y]
}
