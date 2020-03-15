/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/



package main


// 一共max / 2行
// 每行记得少一个。
// [i][j] [j][max-i-1] [max-i-1][max-j-1] [max-j-1][i]
func rotate(matrix [][]int)  {
    for i := 0 ; i < len(matrix) / 2; i++{
        for j := i; j < len(matrix) - i - 1; j ++{
            matrix[i][j],matrix[j][len(matrix) - i - 1],matrix[len(matrix) - i - 1][len(matrix) - j - 1],matrix[len(matrix) - j - 1][i] = matrix[len(matrix) - j - 1][i], matrix[i][j],matrix[j][len(matrix) - i - 1],matrix[len(matrix) - i - 1][len(matrix) - j - 1]
        }
    }     
}


func rotate(matrix [][]int)  {
    for i := 0 ; i < len(matrix) / 2; i++{
        for j := i; j < len(matrix) - i - 1; j ++{
            tmp := matrix[i][j]
            matrix[i][j] = matrix[len(matrix) - j - 1][i]
            matrix[len(matrix) - j - 1][i] = matrix[len(matrix) - i - 1][len(matrix) - j - 1]
            matrix[len(matrix) - i - 1][len(matrix) - j - 1] = matrix[j][len(matrix) - i - 1]
            matrix[j][len(matrix) - i - 1] = tmp
        }
    }     
}