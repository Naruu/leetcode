func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for idx, num := range nums {
        if pair_idx, ok := m[target - num]; ok {
            return []int{pair_idx, idx}
        }
        m[num] = idx
    }
    return []int{-1, -1}
}
