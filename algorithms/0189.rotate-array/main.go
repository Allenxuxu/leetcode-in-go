package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 将最后 k 个元素暂存
// 将前面的元素逐个后移动 k 位
// 将暂存的 k 个元素复制的数组前面
func rotate1(nums []int, k int) {
	k = k % len(nums)
	tmp := make([]int, k)
	copy(tmp, nums[len(nums)-k:])

	toMoveLen := len(nums) - k
	for i := toMoveLen - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
// 假设 n=7 且 k=3 。
// 原始数组                     : 1 2 3 4 5 6 7
// 1. 反转所有数字后             : 7 6 5 4 3 2 1
// 2. 反转前 k 个数字后          : 5 6 7 4 3 2 1
// 3. 反转后 n-k 个数字后        : 5 6 7 1 2 3 4 --> 结果
func rotate2(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

// reverse 反转数组
func reverse(nums []int, i, j int) {
	for i < j {
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
