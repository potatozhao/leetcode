/*
*/

package main

func exist(board [][]byte, word string) bool {
    x := len(board)
    if x <=0{
        return false
    }
    y := len(board[0])
    if x * y < len(word){
        return false
    }
    markedMap := make([][]int, x)
    for i := range markedMap{
        markedMap[i] = make([]int, y)
    }
    // markedList := make([]map[string]bool, len(word))
    // for i := range marketList{
        // markedList[i] = make(map[string]bool)
    // }

    for i := 0;i < x;i ++ {
        for j := 0; j< y; j++{
            if dfs(i,j,0,word,board,markedMap){
                return true
            }
        }
    }
    return false
}

func dfs(x,y,pos int, word string, board [][]byte, markedMap [][]int) bool {
    if pos >= len(word){
        return true
    }
    if x <0 || y < 0 || x >= len(board) || y >= len(board[0]) || word[pos] != board[x][y] || markedMap[x][y] == 1{
        return false
    }
    markedMap[x][y] = 1
    if dfs(x+1, y, pos + 1, word, board, markedMap) || dfs(x, y+1, pos + 1, word, board, markedMap) || dfs(x-1, y, pos + 1, word, board, markedMap) || dfs(x, y-1, pos + 1, word, board, markedMap){
        return true
    }
    markedMap[x][y] = 0
    return false
}


// 降低内存使用。
func exist(board [][]byte, word string) bool {
    x := len(board)
    if x <=0{
        return false
    }
    y := len(board[0])
    if x * y < len(word){
        return false
    }

    for i := 0;i < x;i ++ {
        for j := 0; j< y; j++{
            if dfs(i,j,0,word,board){
                return true
            }
        }
    }
    return false
}

func dfs(x,y,pos int, word string, board [][]byte) bool {
    if pos >= len(word){
        return true
    }
    if x <0 || y < 0 || x >= len(board) || y >= len(board[0]) || word[pos] != board[x][y]{
        return false
    }
    board[x][y] = 0
    if dfs(x+1, y, pos + 1, word, board) || dfs(x, y+1, pos + 1, word, board) || dfs(x-1, y, pos + 1, word, board) || dfs(x, y-1, pos + 1, word, board){
        return true
    }
    board[x][y] = word[pos]
    return false
}