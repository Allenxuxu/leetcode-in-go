package algorithm

var m = map[uint8]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var nums []string
	for i := 0; i < len(digits); i++ {
		nums = append(nums, m[digits[i]-'0'])
	}

	var ret []string
	helper(0, nums, "", &ret)
	return ret
}

func helper(index int, nums []string, ans string, result *[]string) {
	if len(ans) == len(nums) {
		*result = append(*result, ans)
		return
	}

	for _, v := range nums[index] {
		helper(index+1, nums, ans+string(v), result)
	}
}
