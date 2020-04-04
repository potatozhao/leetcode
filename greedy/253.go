/*
253. 会议室 II
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

示例 1:

输入: [[0, 30],[5, 10],[15, 20]]
输出: 2
示例 2:

输入: [[7,10],[2,4]]
输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/meeting-rooms-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


package main

func max(a,b int) int {
    if a >b {
        return a
    }
    return b
}

// 暴力法
func minMeetingRooms(intervals [][]int) int {
    maxNum := 0
    for i := range intervals{
        maxNum = max(intervals[i][1], maxNum)
    }
    tmpList := make([]int, maxNum+1)
    for i := range intervals{
        for j := intervals[i][0]; j < intervals[i][1]; j++{
            tmpList[j]++
        }
    }
    maxNum = 0
    for i := range tmpList{
        maxNum = max(maxNum, tmpList[i])
    }
    return maxNum
}


// 统计学方法，把开始和结束时间放一起，排序，开始和结束时间相同时，结束放在前面，遍历，开始++，结束--，求过程中最大值。
type Intslice [][2]int

func (p Intslice) Len() int           { return len(p) }
func (p Intslice) Less(i, j int) bool {
    return (p[i][0] < p[j][0]) || (p[i][0] == p[j][0] && p[i][1]<p[j][1])
}
func (p Intslice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func minMeetingRooms(intervals [][]int) int {
    tmpList := make([][2]int, 0,len(intervals) * 2)
    for i := range intervals{
        tmpList = append(tmpList, [2]int{intervals[i][0], 1})
        tmpList = append(tmpList, [2]int{intervals[i][1], -1})
    }
    sort.Sort(Intslice(tmpList))
    maxNum := 0
    tmpNum := 0
    for i := range tmpList{
        if tmpList[i][1] < 0{
            tmpNum--
        }else{
            tmpNum++
        }
        maxNum = max(maxNum, tmpNum)
    }
    return maxNum
}