/*
对于一个具有树特征的无向图，我们可选择任何一个节点作为根。图因此可以成为树，在所有可能的树中，具有最小高度的树被称为最小高度树。给出这样的一个图，写出一个函数找到所有的最小高度树并返回他们的根节点。

格式

该图包含 n 个节点，标记为 0 到 n - 1。给定数字 n 和一个无向边 edges 列表（每一个边都是一对标签）。

你可以假设没有重复的边会出现在 edges 中。由于所有的边都是无向边， [0, 1]和 [1, 0] 是相同的，因此不会同时出现在 edges 里。

示例 1:

输入: n = 4, edges = [[1, 0], [1, 2], [1, 3]]

        0
        |
        1
       / \
      2   3

输出: [1]
示例 2:

输入: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]

     0  1  2
      \ | /
        3
        |
        4
        |
        5

输出: [3, 4]
说明:

 根据树的定义，树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
树的高度是指根节点和叶子节点之间最长向下路径上边的数量。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-height-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// dfs 方案，TLE了。。。。。
func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	intMax := int(^uint(0) >> 1)
	graphMap := make(map[int][]int)
	for _, v := range edges {
		for i := range v {
			if _, ok := graphMap[v[i]]; !ok {
				graphMap[v[i]] = make([]int, 0)
			}
		}
		graphMap[v[0]] = append(graphMap[v[0]], v[1])
		graphMap[v[1]] = append(graphMap[v[1]], v[0])
	}
	maxNum := intMax
	maxList := make([]int, 0)
	for k, _ := range graphMap {
		tmp := dfs(graphMap, -1, k)
		if tmp == maxNum {
			maxList = append(maxList, k)
		}
		if tmp < maxNum {
			maxList = make([]int, 0)
			maxNum = tmp
			maxList = append(maxList, k)
		}
	}
	return maxList
}

func dfs(graphMap map[int][]int, start, now int) int {
	v, _ := graphMap[now]
	maxNum := 0
	for i := range v {
		if v[i] != start {
			maxNum = max(maxNum, dfs(graphMap, now, v[i]))
		}
	}
	return maxNum + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//bfs 方案、
