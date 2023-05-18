package array

import (
	"math/rand"
	"sort"
	"time"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement_Random(nums []int) int {
	//因为超过n/2的数组下标被众数占据了，这样我们随机挑选一个下标对应的元素并验证，有很大的概率能找到众数
	//由于一个给定的下标对应的数字很有可能是众数，我们随机挑选一个下标，检查它是否是众数，如果是就返回，否则继续随机挑选。
	//生成一个随机数
	//检查随机数对应的元素是否是众数
	//如果是，返回该元素
	//如果不是，重复第一步
	//这个算法的正确性是由题目保证的，因为题目保证给定的数组中一定存在一个众数。
	//时间复杂度：O(N)O(N)
	//空间复杂度：O(1)O(1)

	for {
		rand.Seed(time.Now().UnixNano())
		i := len(nums)
		random := rand.Intn(i - 1)
		if isMajority(nums, nums[random]) {
			return nums[random]
		}
	}
}

func isMajority(nums []int, i int) bool {
	count := 0
	for _, v := range nums {
		if v == i {
			count++
		}
	}
	return count > len(nums)/2
}

func majorityElement_Boyer_Moore(nums []int) int {
	count := 0
	candidate := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
