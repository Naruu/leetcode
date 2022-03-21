func titleToNumber(columnTitle string) int {
    answer := 0
    mul := 1
    for i := len(columnTitle) - 1; i >= 0; i-- {
        answer += (int(columnTitle[i]) - int('A') + 1) * mul
        mul *= 26
    }
    return answer

}
