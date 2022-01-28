func plusOne(digits []int) []int {
    var answer []int
    carry := 1
    for i := len(digits) - 1 ; i >= 0; i -= 1 {
        s := carry + digits[i]
        answer = append(answer, s % 10)
        carry = s / 10
    }
    if carry > 0 {
        answer = append(answer, carry)
    }

    for i := 0; i < len(answer) / 2; i += 1 {
        answer[i], answer[len(answer) - 1 - i] = answer[len(answer) - 1 - i], answer[i]
    }
    return answer

}