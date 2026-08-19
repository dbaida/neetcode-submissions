func removeElement(nums []int, val int) int {
    filteredLength := 0
    for _, num := range nums {
        if num != val {
            nums[filteredLength] = num
            filteredLength++
        }
    }
    nums = nums[:filteredLength]

    return len(nums)
}
