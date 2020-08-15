/*
34. 在排序数组中查找元素的第一个和最后一个位置
难度
中等

549





给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/


func searchRange(nums []int, target int) []int {
    if len(nums) <= 0{
        return []int{-1,-1}
    }
    s := binarySearch(nums, target, true)
    e := binarySearch(nums, target, false)
    return []int{s,e}
}


func binarySearch(list []int,key int, flag bool) int{
    start := 0
    end := len(list) - 1
    mid := -1
    for start <= end{
        mid = (start + end) >> 1
        if flag{
            if list[mid] >= key{
                end = mid -1
            }else{
                start = mid + 1
            }
        }else{
            if list[mid] > key{
                end = mid -1
            }else{
                start = mid +1
            }
        }
    }
    if list[mid] < key{
        mid += 1
    }else if list[mid] > key{
        mid -= 1
    }
    if mid < 0 || mid >= len(list) || list[mid] != key{
        return -1
    }
    return mid
}