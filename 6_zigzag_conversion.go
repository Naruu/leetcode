func convert(s string, numRows int) string {
    groupByRow := func (s string, numRows int) map[int] []byte {
        grouped := map[int] []byte {}
        for idx := 0; ; {
            for i := 0; i < numRows; i += 1 {
                grouped[i] = append(grouped[i], s[idx])
                idx += 1
                if idx >= len(s) {
                    return grouped
                }
            }
            for i := numRows - 2; i > 0; i -= 1 {
                grouped[i] = append(grouped[i], s[idx])
                idx += 1
                if idx >= len(s) {
                    return grouped
                }
            }
        }
    }
    grouped := groupByRow(s, numRows)
    answer := ""
    for i := 0; i < numRows; i+= 1 {
        answer += string(grouped[i])
    }

    return answer

}
