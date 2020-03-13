/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/



// 先按照起始位置排序，然后轮询，判断是否要与前一个合并，使用双指针，双数组的方案来做。
type IntSlice [][]int
 
func (s IntSlice) Len() int { return len(s) }
 
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
 
func (s IntSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }

func merge(intervals [][]int) [][]int {
    if len(intervals)<= 1{
        return intervals
    }
    ret := make([][]int, 0)
    sort.Sort(IntSlice(intervals))
    ret = append(ret, intervals[0]) // 第一个先入数组
    for i := range intervals{
        s := intervals[i][0]
        e := intervals[i][1]
        if s <= ret[len(ret) - 1][1]{ // 判断是否需要合并，因为排序过了，不需要考虑s过小的问题。
            ret[len(ret) - 1][1] = max(ret[len(ret) - 1][1], e) // 判断区间结尾
        }else{
            ret = append(ret, intervals[i])
        }
    }
    return ret
}

func max(a,b int)int{
    if a > b{
        return a
    }
    return b
}

// 内存优化版本
func merge2(intervals [][]int) [][]int {
    if len(intervals)<= 1{
        return intervals
    }
    sort.Sort(IntSlice(intervals))
    a := 0
    for i := range intervals{
        if intervals[i][0] <= intervals[a][1]{
            intervals[a][1] = max(intervals[a][1], intervals[i][1])
        }else{
            a += 1
            intervals[a][0],intervals[a][1]  = intervals[i][0],intervals[i][1]
        }
    }
    return intervals[:a+1]
}