package algorithm

import (
	"math"
)

//通过循环将数字x的每一位拆开，在计算新值时每一步都判断是否溢出。
//溢出条件有两个，一个是大于整数最大值MAX_VALUE，另一个是小于整数最小值MIN_VALUE，设当前计算结果为ans，下一位为pop。
//从ans * 10 + pop > MAX_VALUE这个溢出条件来看
//当出现 ans > MAX_VALUE / 10 且 还有pop需要添加 时，则一定溢出
//当出现 ans == MAX_VALUE / 10 且 pop > 7 时，则一定溢出，7是 2^31 - 1 (2147483647) 的个位数
//从ans * 10 + pop < MIN_VALUE这个溢出条件来看
//当出现 ans < MIN_VALUE / 10 且 还有pop需要添加 时，则一定溢出
//当出现 ans == MIN_VALUE / 10 且 pop < -8 时，则一定溢出，8是 -2^31(-2147483648) 的个位数
// -2147483648
// 2147483647
func reverse(x int) int {
	res := 0
	for x != 0 {
		// 取出x的末尾
		temp := x % 10
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && temp > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && temp < -8) {
			return 0
		}
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	return res
}
