package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(1)
// 双指针
// i，j 两个指针 i 永远小于 j
// 当 nums[i]  == 0 时候，自增 j 直到找到 nums[j] != 0 ,然后交换两个值
// 不停的将 0 往后交换，最后所有的 0 都在数组末尾
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
// 遍历数组，如果数组元素不为 0 ，则逐个将元素复制得到数组开头
// 遍历完成后所有的 非 0 元素都在数组开头
// 通过 index 来记录，最后一个 非0 元素下标，将 index 之后的元素置 0
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

func moveZeroes2(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func moveZeroes3(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			nums[i-count], nums[i] = nums[i], nums[i-count]
		}
	}
}
