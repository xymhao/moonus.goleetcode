package hard

//There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
//
//You are giving candies to these children subjected to the following requirements:
//
//Each child must have at least one candy.
//Children with a higher rating get more candies than their neighbors.
//Return the minimum number of candies you need to have to distribute the candies to the children.
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/candy

func candy(ratings []int) int {
	var result = 0
	//	1. 从左到右遍历，如果右边的评分比左边的高，那么右边的糖果数等于左边的糖果数+1
	left := make([]int, len(ratings))
	n := len(ratings)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	//	2. 从右到左遍历，如果左边的评分比右边的高，那么左边的糖果数等于右边的糖果数+1
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		result += max(left[i], right)
	}
	return result
}

func max(i int, right int) int {
	if i > right {
		return i
	}
	return right
}
