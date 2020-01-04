import "fmt"

type TicTacToe struct {
    N int
    maps map[string]int
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    ret := TicTacToe{}
    ret.maps = make(map[string]int)
    ret.N = n
    return ret
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    this.maps[fmt.Sprintf("%d_l_%d", player, row)] += 1
    this.maps[fmt.Sprintf("%d_r_%d", player, col)] += 1
    this.maps[fmt.Sprintf("%d_lr_%d", player, row+col)] += 1
    this.maps[fmt.Sprintf("%d_rl_%d", player, row-col)] += 1
    if this.maps[fmt.Sprintf("%d_l_%d", player, row)] == this.N{
        return player
    }
    if this.maps[fmt.Sprintf("%d_r_%d", player, col)] == this.N{
        return player
    }
    if this.maps[fmt.Sprintf("%d_lr_%d", player, row+col)] == this.N{
        return player
    }
    if this.maps[fmt.Sprintf("%d_rl_%d", player, row-col)] == this.N{
        return player
    }
    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */