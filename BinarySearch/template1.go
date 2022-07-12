package BinarySearch

func BinarySearch(nums []int, target int)  int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid -1
		}
	}
}


// 二分查找的最基础模式，无需后续处理能直接返回所查找的元素