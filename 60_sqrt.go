func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }

    left := 0
    right := x
    mid := 0
    for (right - left) > 1 {
        mid = (left + right) /2
        square := mid * mid
        if square == x {
            left = mid
            break
            // answer is mid
        }
        if square < x {
            left = mid
        } else {
            right = mid
        }
    }
    return left
}