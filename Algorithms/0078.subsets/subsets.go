package problem0078

<<<<<<< HEAD
import (
	"math"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0, resSize(nums))
	rec(nums, []int{}, &res)
	return res
}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func rec(nums, temp []int, res *[][]int) {
	size := len(nums)
	if size == 0 {
		*res = append(*res, temp)
		return
	}

	// 对于 last 来说，要么存在于某个子集中，要么不存在
	nums, last := nums[:size-1], nums[size-1]

	rec(nums, temp, res)

	withLast := make([]int, 1, 1+len(temp))
	withLast[0] = last
	withLast = append(withLast, temp...)
	rec(nums, withLast, res)
=======
func subsets(nums []int) [][]int {
	res := make([][]int, 1, 1024)
	for _, n := range nums {
		for _, r := range res {
			res = append(res, append([]int{n}, r...))
		}
	}
	return res
>>>>>>> f33c3a477711033e1c5c5c04e72ce2c3c83f449e
}
