/*
假设你是一位顺风车司机，车上最初有 capacity 个空座位可以用来载客。由于道路的限制，车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向，你可以将其想象为一个向量）。

这儿有一份行程计划表 trips[][]，其中 trips[i] = [num_passengers, start_location, end_location] 包含了你的第 i 次行程信息：

必须接送的乘客数量；
乘客的上车地点；
以及乘客的下车地点。
这些给出的地点位置是从你的 初始 出发位置向前行驶到这些地点所需的距离（它们一定在你的行驶方向上）。

请你根据给出的行程计划表和车子的座位数，来判断你的车是否可以顺利完成接送所用乘客的任务（当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false）。

 

示例 1：

输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
示例 2：

输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true
示例 3：

输入：trips = [[2,1,5],[3,5,7]], capacity = 3
输出：true
示例 4：

输入：trips = [[3,2,7],[3,7,9],[8,3,9]], capacity = 11
输出：true
 

提示：

你可以假设乘客会自觉遵守 “先下后上” 的良好素质
trips.length <= 1000
trips[i].length == 3
1 <= trips[i][0] <= 100
0 <= trips[i][1] < trips[i][2] <= 1000
1 <= capacity <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/car-pooling
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


package main


// 解法。。。先存储每个站点的下车人数，再把上车站点和瞎扯站点存成list， 然后排序，双指针走一遍。
type IntSlice []int
 
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func carPooling(trips [][]int, capacity int) bool {
    startMap :=make(map[int]int)
    endMap := make(map[int]int)
    nowP := 0
    inList := make([]int, 0,len(trips))
    outList := make([]int, 0,len(trips))
    for i := range trips{
        if _, ok:= startMap[trips[i][1]]; !ok{
            startMap[trips[i][1]] = trips[i][0]
            inList = append(inList, trips[i][1])
        }else{
            startMap[trips[i][1]] += trips[i][0]
        }
        if _, ok:= endMap[trips[i][2]]; !ok{
            endMap[trips[i][2]] = trips[i][0]
            outList = append(outList, trips[i][2])
        }else{
            endMap[trips[i][2]] += trips[i][0]
        }
    }
    sort.Sort(IntSlice(inList))
    sort.Sort(IntSlice(outList))

    inNum := 0
    outNum := 0
    for inNum < len(inList) || outNum < len(outList){
        if inNum < len(inList) && outNum < len(outList){
            if inList[inNum] < outList[outNum]{
                nowP += startMap[inList[inNum]]
                inNum++
            }else if inList[inNum] > outList[outNum]{
                nowP -= endMap[outList[outNum]]
                outNum++
            }else{
                nowP += (startMap[inList[inNum]] - endMap[outList[outNum]])
                inNum++
                outNum++
            }
            if nowP > capacity{
                return false
            }
        }else{
            break
        }
    }
    return true
}