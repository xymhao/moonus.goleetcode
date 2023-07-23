package array

import "math/rand"

type RandomizedSet struct {
	//O(1) 时间范围内获取数据
	nums []int
	//O(1) 时间范围内判断数据是否存在
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{nums: []int{}, indices: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indices[val]
	if ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.indices[val]
	if !ok {
		return false
	}
	//将要删除的元素和最后一个元素交换位置
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last]
	//更新索引未知
	this.indices[this.nums[id]] = id
	delete(this.indices, val)
	this.nums = this.nums[:last]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	//生成随机数
	i := len(this.nums)
	if i == 0 {
		return 0
	}
	randomVal := rand.Intn(i)
	return this.nums[randomVal]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
