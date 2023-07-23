package array

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i, _ := range nums {
		val := 1
		for j := 0; j < len(nums); j++ {
			if j != i {
				val *= nums[j]
			}
		}
		result[i] = val
	}
	return result
}

//	1,2,3,4
//
// 前缀之积L：1,1,2,6
// 后缀之积R：24,12,4,1
// 结果：L[i-1]*R[i+1]
//
//	24，12，8，6
func productExceptSelf2(nums []int) []int {
	l := make([]int, len(nums))
	R := make([]int, len(nums))
	result := make([]int, len(nums))

	l[0] = 1
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			l[i] = l[i-1] * nums[i-1]
		}
	}

	R[len(nums)-1] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 {
			R[i] = R[i+1] * nums[i+1]
		}
	}

	for i := 0; i < len(nums); i++ {
		result[i] = l[i] * R[i]
	}

	return result
}
