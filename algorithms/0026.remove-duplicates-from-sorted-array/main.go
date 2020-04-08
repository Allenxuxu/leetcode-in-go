package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(1)
// 快慢指针 / 双指针
// i 慢指针 j 快指针 , i 永远小于 j
// 如果 nums[i] == nums[j] 说明重复了，
// j++ 快指针后移动，直到找到一个和 nums[i]，nums[i+1] = nums[j]，在 i++ j++ 同时后移
//
// 会有这样一种情况，如 [1,2,3] ，nums[i+1] = nums[j] 时候， i+1 其实一直是等于 j 的， 可以判断 if i+1 != j 来避免重复拷贝。
func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1
}
