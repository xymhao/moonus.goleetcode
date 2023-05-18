package array

func removeElement(nums []int, val int) int {
	//双指针
	//left：表示要写入人的位置
	//right：表示当前遍历到的位置，过滤掉val， 写入left
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}
