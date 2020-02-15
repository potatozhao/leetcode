/*
现在你总共有 n 门课需要选，记为 0 到 n-1。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？

示例 1:

输入: 2, [[1,0]] 
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:

输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
说明:

输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
提示:

这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

// 这是个有向图题，判断是否成环

// UnionFind不适用有向图。
type UnionFind struct{
    Prev []int
}

func NewUnionFind(l int) UnionFind{
    uf := UnionFind{
        Prev: make([]int ,l),
    }
    for i := range uf.Prev{
        uf.Prev[i] = i
    }
    return uf
}

func (this *UnionFind)Find(a int) int{
    pa := this.Prev[a]
    if pa == a{
        return pa
    }
    return this.Find(pa)
}

func (this *UnionFind)Union(a,b int){
    pa := this.Find(a)
    pb := this.Find(b)
    // 已经链接了
    if pa == pb {
        return
    }
    //有向图链接
    this.Prev[pa] = pb
    return
}

func (this *UnionFind)CheckRing(a, n int)int{
    if this.Prev[a] == a{
        return a
    }
    if n >= len(this.Prev){
        return -1
    }
    return this.CheckRing(this.Prev[a], n+1)
}

// dfs解法。没有剪枝，请注意思考。
func canFinish(numCourses int, prerequisites [][]int) bool {
    drawMap := make(map[int][]int)
    for _,l := range(prerequisites){
        if _, ok := drawMap[l[0]]; !ok{
            drawMap[l[0]]= make([]int, 0)
        }
        drawMap[l[0]] = append(drawMap[l[0]], l[1])
    }
    markList := make([]int,numCourses)
    for k := range drawMap{
        markList[k] = 1
        if dfs(drawMap, k, markList) {
            return false
        }
        markList[k] = 2
    }
    return true
}

func dfs(drawMap map[int][]int, k int, markList []int) bool {
    pkList, ok:= drawMap[k]
    if !ok{
        return false
    }
    for _,v := range pkList{
        if markList[v] ==1 {
            return true
        }
        if markList[v] == 2{
            continue
        }
        markList[v] = 1
        if dfs(drawMap, v, markList){
            return true
        }
        markList[v] = 2
    }
    return false
}