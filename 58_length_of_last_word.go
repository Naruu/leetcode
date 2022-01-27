func lengthOfLastWord(s string) int {
    splited := strings.Fields(s)
    return len(splited[len(splited)-1])
}