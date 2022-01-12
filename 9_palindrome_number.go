func isPalindrome1(x int) bool {
    if x < 0 {
        return false
    }

    var nums []int
    c := 0
    for q := x; q > 0 ; q = q / 10 {
        nums = append(nums, q % 10)
        c += 1
    }

    for i, q := 0, x; q > 0 && i < c / 2; q, i = q / 10, i + 1 {
        if q % 10 != nums[c - i -1] {
            return false
        }
    }
    return true
}

func isPalindrome2(x int) bool {
    num := strconv.Itoa(x)
    for i := 0; i< len(num) / 2; i++ {
        if num[i] != num[len(num) - i - 1] {
            return false
        }
    }
    return true
}