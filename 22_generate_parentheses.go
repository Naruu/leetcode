// yet failed
import (
    "fmt"
)

func combi(n int, cache map[int][][]int) [][]int {
    if v, err := cache[n]; err == true {
        return v;
    }

    var answer [][]int
    for i := 1; i < n; i += 1 {
        for _, c := range(combi(n -i, cache)) {
            one := append(c, i)
            answer = append(answer, one)
        }
    }
    answer = append(answer, []int{n})
    cache[n] = answer
    return answer
}

func generateParenthesis(n int) []string {
    palindromeMap := map[int]string {
        1: "()",
        2: "(())",
        3: "((()))",
        4: "(((())))",
        5: "((((()))))",
        6: "(((((())))))",
        7: "((((((()))))))",
        8: "(((((((())))))))",
    }

    cache := map[int][][]int{}
    combi(n, cache)
    answer := []string{}
    fmt.Println(cache[n])
    for _, c := range(cache[n]) {
        one := ""
        for _, i := range(c) {
            one += palindromeMap[i]
        }
        answer = append(answer, one)
    }
    return answer
}