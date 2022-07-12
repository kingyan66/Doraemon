package BinarySearch

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left + 1 < right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	// 后续处理
}

// 实现二分的另一种方法，常考
// 搜索条件需要访问元素的直接左右邻居
// 使用元素的邻居来决定向左还是向右
// 保证查找空间在每个步骤中至少还有三个元素
// 需要进行后处理，当剩下2个元素时，循环结束。最后评估其余元素是否符合条件