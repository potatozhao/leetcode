/*
编写一个程序，找出第 n 个丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。

*/

package main

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	num := 1
	numList := make([]int, n)
	numList[0] = 1
	a := 0
	b := 0
	c := 0
	for num < n {
		tmp := min(2*numList[a], 3*numList[b], 5*numList[c])
		if tmp == 2*numList[a] {
			a++
		}
		if tmp == 3*numList[b] {
			b++
		}
		if tmp == 5*numList[c] {
			c++
		}
		numList[num] = tmp
		num++
	}
	return numList[n-1]
}

func min(a, b, c int) int {
	if b <= a && b <= c {
		return baa
	} else if c <= a && c <= b {
		return c
	} else {
		return a
	}
}

// 解法很trick。。。。题目不好
