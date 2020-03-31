/*

假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。

注意：
总人数少于1100人。

示例

输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


/*
假设候选队列为 A，已经站好队的队列为 B.

从 A 里挑身高最高的人 x 出来，插入到 B. 因为 B 中每个人的身高都比 x 要高，因此 x 插入的位置，就是看 x 前面应该有多少人就行了。比如 x 前面有 5 个人，那 x 就插入到队列 B 的第 5 个位置。

作者：ivan_allen
链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height/solution/406-gen-ju-shen-gao-zhong-jian-dui-lie-pai-xu-hou-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package main


type ListSlice [][]int

func (p ListSlice) Len() int           { return len(p) }
func (p ListSlice) Less(i, j int) bool {
    if p[i][0] > p[j][0]{
        return true
    }else if p[i][0] == p[j][0] && p[i][1] < p[j][1]{
        return true
    }
    return false
}
func (p ListSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reconstructQueue(people [][]int) [][]int {
    sort.Sort(ListSlice(people))
    ret := make([][]int, len(people))
    for i := range people{
        start := people[i][1]
        for j := len(ret) - 1; j > start;j--{
            ret[j] = ret[j-1]
        }
        ret[start] = people[i]
    }
    return ret
}