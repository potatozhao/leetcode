/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

 

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
 

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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