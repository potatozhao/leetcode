/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func numIslands(grid [][]byte) int {
    if len(grid) <= 0{
        return 0
    }
    dfsMap := make([][]byte, len(grid))
    for i := range dfsMap{
        dfsMap[i] = make([]byte,len(grid[0]))
    }

    num := 0

    for i := range grid{
        for j := range grid[i]{
            num += dfs(grid, i,j,dfsMap)
        }
    }

    return num 
}

func dfs(grid [][]byte, i,j int, dfsMap [][]byte)int{
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' ||dfsMap[i][j] == 1{
        return 0
    }
    dfsMap[i][j] = 1
    dfs(grid, i+1, j, dfsMap)
    dfs(grid, i-1, j, dfsMap)
    dfs(grid, i, j+1, dfsMap)
    dfs(grid, i, j-1, dfsMap)
    return 1
}

// 使用本身矩阵当剪枝矩阵
func numIslands2(grid [][]byte) int {
    if len(grid) <= 0{
        return 0
    }

    num := 0

    for i := range grid{
        for j := range grid[i]{
            num += dfs2(grid, i,j)
        }
    }

    return num 
}

func dfs2(grid [][]byte, i,j int)int{
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0'{
        return 0
    }
    grid[i][j] = '0'
    dfs2(grid, i+1, j)
    dfs2(grid, i-1, j)
    dfs2(grid, i, j+1)
    dfs2(grid, i, j-1)
    return 1
}

// bfs解法
type Node struct{
    X int
    Y int
}

func bfs(grid [][]byte, i,j int)int{
    if grid[i][j] < '1'{
        return 0
    }
    queue := make([]Node, 0)
    grid[i][j] = 0
    queue = append(queue, Node{X:i,Y:j})
    for len(queue) > 0{
        tmpNode := queue[0]
        queue = queue[1:]
        x := tmpNode.X
        y := tmpNode.Y
        if x - 1 >= 0 && grid[x-1][y] == '1'{
            grid[x-1][y] = 0
            queue = append(queue,Node{X:x-1,Y:y})
        }
        if x + 1 < len(grid) && grid[x+1][y] == '1'{
            grid[x+1][y] = 0
            queue = append(queue,Node{X:x+1,Y:y})
        }
        if y - 1 >= 0 && grid[x][y-1]  == '1'{
            grid[x][y-1] = 0
            queue = append(queue,Node{X:x,Y:y-1})
        }
        if y + 1 < len(grid[0]) && grid[x][y+1]  == '1'{
            grid[x][y+1] = 0
            queue = append(queue,Node{X:x,Y:y+1})
        }
    }
    return 1
}