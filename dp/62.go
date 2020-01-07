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