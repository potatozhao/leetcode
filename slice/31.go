/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/


//就是整一个比当前数字表示的数更大一点的最小数。
// 找到最后一个升序的位置，即2，3，7，6，1中的3位置，然后从3之后的数据，取出比3大的最小值6，然后互换位置，对6后面的数值进行从小到大排列。
package main

func nextPermutation(nums []int)  {
    if len(nums) <=1{
        return
    }
    pos := -1
    min := -1
    for i := 0; i < len(nums); i++{
        if i < len(nums) - 1 && nums[i] < nums[i+1]{
            pos = i
            min = -1
        }else{
            if min == -1 || (nums[min] >= nums[i]&& pos >= 0 && nums[i] > nums[pos]){
                min = i
            }
        }
    }
    if pos < 0{
        sort.Ints(nums)
	}else{
		nums[pos],nums[min] = nums[min], nums[pos]
        sort.Ints(nums[pos+1:])
	}
	return
}