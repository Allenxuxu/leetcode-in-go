package algorithm

func maxArea(height []int) int {
	var ret int
	for i, j := 0, len(height)-1; i < j; {
		var area int
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}

		if area > ret {
			ret = area
		}
	}

	return ret
}
