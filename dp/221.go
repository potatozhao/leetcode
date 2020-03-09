/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix[0]) == 0 {
		return 0
	}

	tmpV := make([][]byte, len(matrix))
	for i, _ := range tmpV {
		tmpV[i] = make([]byte, len(matrix[0]))
	}

	var maxNum byte

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if (i == 0 || j == 0) && matrix[i][j] == '1' {
				tmpV[i][j] = 1
			} else if matrix[i][j] == '1' {
				tmpV[i][j] = min(tmpV[i-1][j-1], tmpV[i][j-1], tmpV[i-1][j]) + 1
			}
			if tmpV[i][j] > maxNum {
				maxNum = tmpV[i][j]
			}
		}
	}
	return int(maxNum) * int(maxNum)
}

func min(a, b, c byte) byte {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
