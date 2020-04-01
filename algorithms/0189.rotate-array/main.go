package algorithm

// rotate1 解法1
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

// rotate2 解法2
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
