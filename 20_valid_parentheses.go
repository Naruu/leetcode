func isValid(s string) bool {
    PAIR := map[rune]rune {
        '(': ')',
        '{': '}',
        '[': ']',
    }
    var stack []rune

    for _, str := range s {
        if str == '(' || str == '{' || str == '[' {
            stack = append(stack, str)
            continue
        }
        if len(stack) < 1 {
            return false;
        }
        if str != PAIR[stack[len(stack) - 1]] {
            return false
        }
        stack = stack[:len(stack) - 1]
    }
    return len(stack) == 0

}