/*
62. 不同路径
难度
中等

641





一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？


*/

func uniquePaths(m int, n int) int {
    if m == 0 || n ==0{
        return 0
    } 
    tmpV := make([][]int, m)
    for i,_ := range tmpV{
        tmpV[i] = make([]int, n)
    }
    for i := 0; i< len(tmpV); i++ {
        for j := 0;j < len(tmpV[i]); j++{
            if (i == 0 || j == 0){
                tmpV[i][j] = 1
            }else{
                tmpV[i][j] = tmpV[i-1][j] + tmpV[i][j - 1]
            }
        }
    }
    return tmpV[m - 1][n - 1]
}