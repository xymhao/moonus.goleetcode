package array

func canJump(nums []int) bool {
	len := len(nums)
	rightMost := 0
	for i := 0; i < len; i++ {
		if i <= rightMost {
			rightMost = Max(rightMost, nums[i]+i)
			if rightMost >= len-1 {
				return true
			}
		}
	}
	return false
}

func Max(most int, i int) int {
	if most > i {
		return most
	}
	return i
}

func jump(nums []int) int {
	//反向查找出发位置
	//1.从最后一个位置往前找，找到能跳到最后一个位置的索引
	position := len(nums) - 1
	steps := 0
	//2.如果有多个位置可以满足，我们可以选择举例最后最远的那个位置
	for position > 0 {
		for i := 0; i < position; i++ {
			if nums[i] >= position-i {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

func jump2(nums []int) int {
	//正向查找可到达的最大位置
	//1.从起始位置开始，找到可到达的最大位置
	//2,3,1,2,4,2,3
	//1.从0出发，能到达1，2，索引位置，索引1是3，索引2 是，索引1能到达更远，因此选择索引1
	steps := 0
	maxPosition := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = Max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
