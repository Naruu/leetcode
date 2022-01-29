func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    answer := make([]byte, len(s))
    ai := 0
	for row := 0; row < numRows ; row += 1 {
        descending := false
        for idx := row ; idx < len(s); {
            step := 2 * (numRows - 1 - row)
            if descending {
                step = 2 * row
            }
            descending = !descending
            if step == 0 {
                continue
            }
            answer[ai] = s[idx]
            idx += step
            ai += 1
        }
	}
    return string(answer)

}
