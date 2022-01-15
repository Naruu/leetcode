SYMBOL_INT_MAP = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
}


class Solution:
    def romanToInt(self, s: str) -> int:
        answer = 0
        for i in range(len(s)-1):
            cur_val = SYMBOL_INT_MAP[s[i]]
            next_val = SYMBOL_INT_MAP[s[i+1]]
            if cur_val >= next_val:
                answer += cur_val
            else:
                answer -= cur_val
        answer += SYMBOL_INT_MAP[s[len(s)-1]]
        return answer
