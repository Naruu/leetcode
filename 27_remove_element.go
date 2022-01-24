func removeElement(nums []int, val int) int {
    idx := 0
    for i := 0; i< len(nums); i += 1 {
        if nums[i] == val {
            continue
        }
        nums[idx] = nums[i]
        idx += 1
    }
    return idx

}