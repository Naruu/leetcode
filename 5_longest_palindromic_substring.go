func longestPalindrome(s string) string {
    answer := ""
    for i := 0; i< len(s); i+= 1 {
        left, right := expandedPalindrome(s, i, i)
        if len(answer) < (right - left) {
            answer = s[left:right]
        }
    }

    for i := 0; i< len(s) - 1; i+= 1 {
        if s[i] == s[i+1] {
            left, right := expandedPalindrome(s, i, i+1)
            if len(answer) < (right - left) {
                answer = s[left:right]
            }
        }
    }
    return answer
}

func expandedPalindrome (str string, left int, right int) (int, int) {
    start, end := left, right
    for (start > 0) && (end < len(str) - 1) {
        if (str[start-1] != str[end+1]) {
            return start, end + 1
        }
        start -= 1
        end += 1
    }
    return start, end + 1
}
