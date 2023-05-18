package array

func removeDuplicates(nums []int) int {
	//创建一个map，key为nums的值，value为nums的下标
	//遍历nums，如果map中不存在当前值，则写入map，同时写入nums
	//如果map中存在当前值，则跳过
	//最后返回nums的长度
	m := make(map[int]int)
	left := 0
	for right := 0; right < len(nums); right++ {
		if _, ok := m[nums[right]]; !ok {
			m[nums[right]] = right
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func removeDuplicatesII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	//slow 和 fast 分别为慢指针和快指针
	//slow 表示处理出的数组长度
	//nums[slow−2]=nums[fast] 时，当前待检查元素 nums[slow -2] == nums[fast] ,nums[fast]不应该被保留
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
