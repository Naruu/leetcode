func maxSubArray(nums []int) int {
    getMax := func (a int, b int ) int {
        if a < b {
            return b
        }
        return a
    }
    curSum := 0
    bestSum := nums[0]
    for i :=0; i < len(nums); i += 1 {
        curSum = getMax(nums[i], curSum + nums[i])
        bestSum = getMax(bestSum, curSum)
    }
    return bestSum

}