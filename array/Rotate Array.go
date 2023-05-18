package array

// rotate 旋转数组 当前下标等于 (当前下标 + k) % 数组长度
// nums = [1,2,3,4,5,6,7], k = 3
//   - -
//     [5,6,7,1,2,3,4]
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, newNums)
}

// //rotate2
func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n

	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}

}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 1 2 3 4 5 6 7
// step1 : 翻转所有元素
// step2： 翻转[0, k mod n - 1]的元素
// step3 : 翻转[k mod n, n - 1]的元素
func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for star, end := 0, len(nums)-1; star < len(nums)/2; star++ {
		nums[star], nums[end-star] = nums[end-star], nums[star]
	}
}
