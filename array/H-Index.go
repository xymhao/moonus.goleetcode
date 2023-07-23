package array

import "sort"

// 3,0,6,1,5
func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	//0,1,3,5,6
	//找到citation[i] > h,说明我们找到了至少被引用了h+1次的论文
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}

func hIndex2(citations []int) int {
	h := 0
	//1.申请一个 len(citations)+1 的数组，用来记录每个引用次数的论文数量
	counter := make([]int, len(citations)+1)

	//2.遍历citations，记录引用次数，如果大于n,则记录在n的位置
	for _, citation := range citations {
		if citation >= len(citations) {
			counter[len(citations)]++
		} else {
			counter[citation]++
		}
	}

	//3.从后往前遍历counter，如果引用次数大于等于h，则说明我们找到了至少被引用了h次的论文
	for i := len(counter) - 1; i >= 0; i-- {
		h += counter[i]
		if h >= i {
			return i
		}
	}
	return 0
}
