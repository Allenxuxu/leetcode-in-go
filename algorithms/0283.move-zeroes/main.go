package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func moveZeroes(nums []int) {
	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		if nums[i] == 0 {
			if nums[j] == 0 {
				j++
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func moveZeroes1(nums []int) {
	var (
		index int
		i     int
	)
	for i, index = 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for j := index; j < len(nums); j++ {
		nums[j] = 0
	}
}
