func climbStairs(n int) int {
    var cache = make([]int, n + 1)
    cache[0] = 1
    cache[1] = 1
    for i := 2; i < n + 1; i+= 1 {
        cache[i] = cache[i-1] + cache[i-2]
    }
    return cache[n]
}