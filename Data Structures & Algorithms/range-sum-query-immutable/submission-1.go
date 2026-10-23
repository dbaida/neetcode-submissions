type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{nums: nums}
}


func (this *NumArray) SumRange(left int, right int) int {
    var result int
	
	for left <= right {
		result += this.nums[left]
		left++

		if left <= right {
			result += this.nums[right]
			right--
		}
	}
	return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */