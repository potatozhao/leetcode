/*

739. 每日温度

根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main


// 比较挫的遍历法
func dailyTemperatures(T []int) []int {
    lastList := make([]int, 101)
    for i := range lastList{
        lastList[i] = 1000000
    }
    maxTmp := 0
    for i := len(T) - 1; i >= 0;i--{
        nowT := T[i]
        if nowT > 100{
            T[i] = 0
        }else if nowT > maxTmp {
            T[i] = 0
            maxTmp = nowT
        }else{
            minNum := 1000000
            for j := nowT + 1; j < len(lastList); j++{
                minNum = min(minNum, lastList[j] - i)
            }
            if minNum < 30000{
                T[i] = minNum
            }else{
                T[i] = 0
            }
        }
        lastList[nowT] = i
    }
    return T
}

func min(a,b int) int {
    if a < b{
        return a
    }
    return b
}

// 栈方案，入栈条件：当前元素比栈顶元素小，出栈条件：遇到比自己大的温度，出栈时索引距离即天数差
func dailyTemperatures(T []int) []int {
    stack := make([]int, 0, 10)
    for i := range T{
        for len(stack) > 0 && T[stack[len(stack)-1]] < T[i]{
            T[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }
    for i := range stack{
        T[stack[i]] = 0
    }
    return T
}