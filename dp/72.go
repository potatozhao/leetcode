/*

72. 编辑距离
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
 

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


/*
首先，两个状态，word1 的i 和 word2的j 
	假如word1[i]==word2[j] 
	那么编辑距离dplist[i][j] = dplist[i-1][j-1]
	假如word1[i]!=word2[j] 
	那么会有三种情况：
		第一种，是word1比word2多了一个字符，那么编辑距离就是删除 dplist[i][j] = dplist[i-1][j] + 1
		第二种，是Word1比word2少了一个字符，那么编辑距离就是增加 dplist[i][j] = dplist[i][j-1] + 1
		第三种，是Word1和word2不多不少，但是要纠错，那么编辑距离就是替换 dplist[i][j] = dplist[i-1][j-1] + 1
	所以动规方程为：
	if word1[i]==word2[j] {
		dplist[i][j] = dplist[i-1][j-1]
	}else{
		dpList[i][j] = min(dpList[i-1][j], dpList[i][j-1], dpList[i-1][j-1]) + 1
	}
	那么矩阵初始化时，base_data:
	需要考虑当word1为空，和word2为空的时候:
	     ‘’ ‘r’ 'o' 's'
	''   0   1   2   3
	'h'  1
	'o'  2
	'r'  3
	's'  4
	'e'  5
*/
package main

func minDistance(word1 string, word2 string) int {
    if len(word1) == 0|| len(word2) == 0{
        return max(len(word1), len(word2))
    }
    dpList := make([][]int, len(word1)+1)
    for i := range dpList{
        dpList[i] = make([]int, len(word2)+1)
    }
    // 初始化矩阵
    dpList[0][0] = 0
    for i := range dpList{
        if i ==0 {
            for j := 1;j< len(dpList[0]); j++{
                dpList[0][j] = j
            }
        }else{
            dpList[i][0] = i
        }
    }
    // 构建dp矩阵
    for i := range word1{
        for j := range word2{
            if word1[i] == word2[j]{
                dpList[i+1][j+1] = dpList[i][j]
            }else{
                dpList[i+1][j+1] = min(dpList[i][j+1], dpList[i+1][j], dpList[i][j]) + 1
            }
        }
    }
    return dpList[len(word1)][len(word2)]
}

func max(a,b int) int {
    if a > b{
        return a
    }
    return b
}

func min(a,b,c int) int {
    if a <= b && a <= c{
        return a
    }
    if b <= c && b <= a{
        return b
    }
    return c
}