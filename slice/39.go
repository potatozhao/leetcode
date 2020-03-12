/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


// 第一，如何去重，第二，剪枝没啥卵用
// 去重，因为是有序数组，所以只要保证不会出现回退到前位就成，保证每次回溯，只能从当前位置进行。
func combinationSum(candidates []int, target int) [][]int {
    return dfs(candidates, target, 0)
}

func dfs(candidates []int, target, pos int)[][]int{
    if target < 0{
        return [][]int{}
    }
    if target == 0{
        return [][]int{{}}
    }
    ret := make([][]int, 0)
    for i := pos; i < len(candidates); i++{
        tmp := dfs(candidates, target - candidates[i], i)
        ret = append(ret, copy(tmp, candidates[i])...)
    }
    return ret
}

func copy(in [][]int, src int) [][]int{
    ret := make([][]int, len(in))
    for i := range in{
        ret[i] = make([]int, 0)
        ret[i] = append(ret[i], in[i]...)
        ret[i] = append(ret[i], src)
    }
    return ret
}
